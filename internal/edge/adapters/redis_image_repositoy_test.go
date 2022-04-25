package adapters

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/redhatinsights/edge-api/config"
)

var (
	redisClient *redis.Client
	redisServer *miniredis.Miniredis
)

// mockRedis creates a miniredis server for testing.
func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return s
}

// setupRedis sets up a redis client for testing.
func setupRedis(t *testing.T) {
	config.Init()
	redisServer = mockRedis()
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
}

// teardownRedis tears down the redis client.
func teardownRedis(t *testing.T) {
	redisClient.Close()
}

func TestNewRedisImageRepository(t *testing.T) {
	setupRedis(t)
	defer teardownRedis(t)

	type args struct {
		db *redis.Client
	}
	tests := []struct {
		name string
		args args
		want *RedisImageRepository
	}{
		{
			name: "should return a new redis image repository",
			args: args{
				db: redisClient,
			},
			want: &RedisImageRepository{
				db: redisClient,
			},
		},
		{
			name: "should panic if db is nil",
			args: args{
				db: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				// recover from panic if one occured.
				if recover() != nil {
					redisClient = nil
				}
			}()
			t.Parallel()
			if got := NewRedisImageRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisImageRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imageKey(t *testing.T) {
	type args struct {
		image *image.Image
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		auth    bool
	}{
		{
			name: "should return the image key",
			args: args{
				image: &validImage,
			},
			want:    fmt.Sprintf("%s:image:%s", common.DefaultAccount.String(), validImage.UUID()),
			wantErr: false,
			auth:    false,
		},
		{
			name: "should fail getting account",
			args: args{
				image: &validImage,
			},
			want:    "",
			wantErr: true,
			auth:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			config.Get().Auth = tt.auth
			got, err := imageKey(tt.args.image)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("imageKey() = %v, want %v", got, tt.want)
			}
			config.Get().Auth = false
		})
	}
}

func TestRedisImageRepository_CreateImage(t *testing.T) {
	setupRedis(t)
	defer teardownRedis(t)
	repository := NewRedisImageRepository(redisClient)

	key, err := imageKey(&validImage)
	if err != nil {
		t.Errorf("imageKey() error = %v", err)
	}

	type args struct {
		ctx   context.Context
		image *image.Image
	}
	tests := []struct {
		name    string
		r       *RedisImageRepository
		args    args
		wantErr bool
		auth    bool
	}{
		{
			name: "should create an image",
			r:    repository,
			args: args{
				ctx:   context.Background(),
				image: &validImage,
			},
			wantErr: false,
			auth:    false,
		},
		{
			name: "should fail create an image, bad account",
			r:    repository,
			args: args{
				ctx:   context.Background(),
				image: &validImage,
			},
			wantErr: true,
			auth:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			config.Get().Auth = tt.auth
			if err := tt.r.CreateImage(tt.args.ctx, tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("RedisImageRepository.CreateImage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if res := redisClient.Get(context.Background(), key); res.Err() != nil {
					t.Errorf("redisClient.Get() error = %s", res.Err())
				}
			}
			config.Get().Auth = false
		})
	}
}

func TestRedisImageRepository_GetImage(t *testing.T) {
	setupRedis(t)
	defer teardownRedis(t)
	repository := NewRedisImageRepository(redisClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("RedisImageRepository.CreateImage() error = %v", err)
	}
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		r       *RedisImageRepository
		args    args
		want    *image.Image
		wantErr bool
		auth    bool
	}{
		{
			name: "should get an image",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
			},
			want:    &validImage,
			wantErr: false,
			auth:    false,
		},
		{
			name: "should fail get an image, bad key",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: uuid.NewString(),
			},
			want:    &validImage,
			wantErr: true,
			auth:    false,
		},
		{
			name: "should fail get an image, bad account",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
			},
			want:    &validImage,
			wantErr: true,
			auth:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			config.Get().Auth = tt.auth
			got, err := tt.r.GetImage(tt.args.ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisImageRepository.GetImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !areEqualImages(got, tt.want) {
				t.Errorf("RedisImageRepository.GetImage() = %v, want %v", got, tt.want)
			}
			config.Get().Auth = false
		})
	}
}

func TestRedisImageRepository_UpdateImage(t *testing.T) {
	setupRedis(t)
	defer teardownRedis(t)
	repository := NewRedisImageRepository(redisClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("RedisImageRepository.CreateImage() error = %v", err)
	}
	type args struct {
		ctx      context.Context
		uuid     string
		updateFn func(image *image.Image) (*image.Image, error)
	}
	tests := []struct {
		name    string
		r       *RedisImageRepository
		args    args
		wantErr bool
	}{
		{
			name: "should update an image",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
				updateFn: func(image *image.Image) (*image.Image, error) {
					image.SetNameAndDesc(common.Name{}, "new description")
					return image, nil
				},
			},
			wantErr: false,
		},
		{
			name: "should fail update an image, bad key",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: uuid.NewString(),
				updateFn: func(image *image.Image) (*image.Image, error) {
					image.SetNameAndDesc(common.Name{}, "new description")
					return image, nil
				},
			},
			wantErr: true,
		},
		{
			name: "should fail update an image, just error",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
				updateFn: func(image *image.Image) (*image.Image, error) {
					image.SetNameAndDesc(common.Name{}, "new description")
					return nil, errors.New("error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UpdateImage(tt.args.ctx, tt.args.uuid, tt.args.updateFn); (err != nil) != tt.wantErr {
				t.Errorf("RedisImageRepository.UpdateImage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				// check if the image has been updated
				got, err := tt.r.GetImage(tt.args.ctx, tt.args.uuid)
				if err != nil {
					t.Errorf("RedisImageRepository.GetImage() error = %v", err)
				}
				if got.Description() != "new description" {
					t.Errorf("RedisImageRepository.UpdateImage() = %v, want new description", got.Description())
				}
			}
		})
	}
}

func TestRedisImageRepository_DeleteImage(t *testing.T) {
	setupRedis(t)
	defer teardownRedis(t)
	repository := NewRedisImageRepository(redisClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("RedisImageRepository.CreateImage() error = %v", err)
	}
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		r       *RedisImageRepository
		args    args
		wantErr bool
		auth    bool
	}{
		{
			name: "should delete an image",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
			},
			wantErr: false,
		},
		{
			name: "should fail delete an image, bad account",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
			},
			wantErr: true,
			auth:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			config.Get().Auth = tt.auth
			if err := tt.r.DeleteImage(tt.args.ctx, tt.args.uuid); (err != nil) != tt.wantErr {
				t.Errorf("RedisImageRepository.DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				// check if the image has been deleted
				_, err := tt.r.GetImage(tt.args.ctx, tt.args.uuid)
				if err == nil {
					t.Errorf("RedisImageRepository.DeleteImage() = %v, want deleted", err)
				}
			}
			config.Get().Auth = false
		})
	}
}

func TestRedisImageRepository_GetImages(t *testing.T) {
	setupRedis(t)
	defer teardownRedis(t)
	repository := NewRedisImageRepository(redisClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("RedisImageRepository.CreateImage() error = %v", err)
	}
	// create the two images
	err = repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("RedisImageRepository.CreateImage() error = %v", err)
	}
	err = repository.CreateImage(context.Background(), &anotherValidImage)
	if err != nil {
		t.Errorf("RedisImageRepository.CreateImage() error = %v", err)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		r       *RedisImageRepository
		args    args
		want    []*image.Image
		wantErr bool
		auth    bool
	}{
		{
			name: "should get all images",
			r:    repository,
			args: args{
				ctx: context.Background(),
			},
			want: []*image.Image{
				&validImage,
				&anotherValidImage,
			},
			wantErr: false,
		},
		{
			name: "should fail get all images, bad account",
			r:    repository,
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			auth:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			config.Get().Auth = tt.auth
			got, err := tt.r.GetImages(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisImageRepository.GetImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotImages := make(map[string]*image.Image, len(got))
			for _, image := range got {
				gotImages[image.UUID()] = image
			}
			for _, image := range tt.want {
				if !areEqualImages(image, gotImages[image.UUID()]) {
					t.Errorf("RedisImageRepository.GetImages() = %v, want %v", image, gotImages[image.UUID()])
				}
			}
			config.Get().Auth = false
		})
	}
}

func TestNewRedisClient(t *testing.T) {
	type args struct {
		cfg *config.EdgeConfig
	}
	tests := []struct {
		name string
		args args
		want *redis.Client
	}{
		{
			name: "should create a redis client",
			args: args{
				cfg: config.Get(),
			},
			want: redis.NewClient(&redis.Options{
				Addr:     ":6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisClient(tt.args.cfg); got.String() != tt.want.String() {
				t.Errorf("NewRedisClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
