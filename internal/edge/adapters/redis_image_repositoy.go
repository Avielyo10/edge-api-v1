package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/go-redis/redis/v8"
	"github.com/redhatinsights/edge-api/config"

	log "github.com/sirupsen/logrus"
)

// RedisImageRepository is a Redis implementation of the Image.Repository interface.
type RedisImageRepository struct {
	db *redis.Client
}

// NewRedisImageRepository returns a new Redis implementation of the Image.Repository interface.
func NewRedisImageRepository(db *redis.Client) *RedisImageRepository {
	if db == nil {
		panic("db cannot be nil")
	}
	return &RedisImageRepository{db: db}
}

// imageKey returns the Redis key for the image.
func imageKey(image *image.Image) (string, error) {
	account, err := image.Account()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s:%s", account.String(), "image", image.UUID()), nil // <- account:image:uuid
}

// CreateImage creates a new image, implementing the Image.Repository interface.
func (r *RedisImageRepository) CreateImage(ctx context.Context, image *image.Image) error {
	log.Debug("redis create image")
	key, err := imageKey(image)
	if err != nil {
		return err
	}
	return r.db.Set(ctx, key, image.MarshalRedis(), 10*time.Minute).Err()
}

// GetImage returns the image with the given UUID, implementing the Image.Repository interface.
func (r *RedisImageRepository) GetImage(ctx context.Context, uuid string) (*image.Image, error) {
	log.WithField("uuid", uuid).Debug("redis get image")
	account, err := common.GetAccountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("%s:%s:%s", account.String(), "image", uuid) // <- account:image:uuid
	result, err := r.db.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	// Unmarshal the result into an image.
	var image image.Image
	err = json.Unmarshal([]byte(result), &image)
	return &image, err
}

// UpdateImage updates the image with the given UUID, implementing the Image.Repository interface.
func (r *RedisImageRepository) UpdateImage(ctx context.Context, uuid string, updateFn func(image *image.Image) (*image.Image, error)) error {
	log.WithField("uuid", uuid).Debug("redis update image")
	image, err := r.GetImage(ctx, uuid)
	if err != nil {
		return err
	}
	updatedImage, err := updateFn(image)
	if err != nil {
		return err
	}
	return r.CreateImage(ctx, updatedImage)
}

// DeleteImage deletes the image with the given UUID, implementing the Image.Repository interface.
func (r *RedisImageRepository) DeleteImage(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("redis delete image")
	account, err := common.GetAccountFromContext(ctx)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s:%s:%s", account.String(), "image", uuid) // <- account:image:uuid
	return r.db.Del(ctx, key).Err()
}

// GetImages returns a list of images, implementing the Image.Repository interface.
func (r *RedisImageRepository) GetImages(ctx context.Context) ([]*image.Image, error) {
	log.Debug("redis get images")
	account, err := common.GetAccountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("%s:%s", account.String(), "image:*") // 000000:image:* <- account:image:*
	iter := r.db.Scan(ctx, 0, prefix, 0).Iterator()
	var images []*image.Image
	for iter.Next(ctx) {
		result, err := r.db.Get(ctx, iter.Val()).Result()
		if err != nil {
			return nil, err
		}
		// Unmarshal the result into an image.
		var image image.Image
		err = json.Unmarshal([]byte(result), &image)
		if err != nil {
			return nil, err
		}
		images = append(images, &image)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return images, nil
}

// NewRedisClient returns a new RedisClient.
func NewRedisClient(cfg *config.EdgeConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
