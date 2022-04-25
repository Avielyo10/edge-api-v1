package adapters

import (
	"context"

	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ReadThroughImageRepository is an implementation of the Image.Repository
type ReadThroughImageRepository struct {
	rdb *RedisImageRepository // redis client
	gdb *GormImageRepository  // gorm client
}

// NewReadThroughImageRepository returns a new ReadThroughImageRepository
func NewReadThroughImageRepository(rdb *redis.Client, gdb *gorm.DB) *ReadThroughImageRepository {
	return &ReadThroughImageRepository{
		rdb: NewRedisImageRepository(rdb),
		gdb: NewGormImageRepository(gdb),
	}
}

// CreateImage creates a new image, implementing the Image.Repository interface.
func (r *ReadThroughImageRepository) CreateImage(ctx context.Context, image *image.Image) error {
	// try to create image in redis
	err := r.rdb.CreateImage(ctx, image)
	if err != nil {
		log.WithField("uuid", image.UUID()).Error(err) // if redis fails, log error
	}
	return r.gdb.CreateImage(ctx, image) // write images synchronously
}

// GetImage returns the image with the given UUID, implementing the Image.Repository interface.
func (r *ReadThroughImageRepository) GetImage(ctx context.Context, uuid string) (*image.Image, error) {
	image, err := r.rdb.GetImage(ctx, uuid) // try to get image from redis
	if err != nil {                         // if redis fails, try gorm
		image, err = r.gdb.GetImage(ctx, uuid) // try to get image from gorm
		if err != nil {                        // if gorm fails, return error
			return nil, err
		}
		// if gorm succeeds, update redis
		go func() {
			err = r.rdb.CreateImage(ctx, image)
			if err != nil { // if redis fails, log error
				log.WithField("uuid", uuid).Error(err)
			}
		}()
	}
	return image, nil
}

// UpdateImage updates the image with the given UUID, implementing the Image.Repository interface.
func (r *ReadThroughImageRepository) UpdateImage(ctx context.Context, uuid string, updateFn func(image *image.Image) (*image.Image, error)) error {
	// try to update image in redis
	err := r.rdb.UpdateImage(ctx, uuid, updateFn)
	if err != nil {
		log.WithField("uuid", uuid).Error(err) // if redis fails, log error
	}
	return r.gdb.UpdateImage(ctx, uuid, updateFn) // try to update image in gorm, synchronously
}

// DeleteImage deletes the image with the given UUID, implementing the Image.Repository interface.
func (r *ReadThroughImageRepository) DeleteImage(ctx context.Context, uuid string) error {
	// try to delete image from redis
	err := r.rdb.DeleteImage(ctx, uuid)
	if err != nil {
		log.WithField("uuid", uuid).Error(err) // if redis fails, log error
	}
	return r.gdb.DeleteImage(ctx, uuid) // try to delete image from gorm, synchronously
}

// GetImages returns a list of images, implementing the Image.Repository interface.
func (r *ReadThroughImageRepository) GetImages(ctx context.Context) ([]*image.Image, error) {
	// we can use redis to help get all images, but this needs to be very suphisticated
	// since each image has an expiration time, and we need to make sure we don't return
	// wrong output.
	// for now, we'll just use gorm
	return r.gdb.GetImages(ctx)
}
