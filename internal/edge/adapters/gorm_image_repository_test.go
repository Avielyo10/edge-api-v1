package adapters

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/common/models"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/google/uuid"
	"github.com/redhatinsights/edge-api/config"
	"gorm.io/gorm"
)

var (
	dbName     string // here so we can tear down the db after each test
	gormClient *gorm.DB
)

// SetupGorm is a helper function to setup a gorm db connection.
func setupGorm(t *testing.T) {
	config.Init()
	dbName = fmt.Sprintf("%s.db", t.Name())
	config.Get().Database.Name = dbName
	gormClient = NewGormClient(config.Get())
}

// TeardownGorm is a helper function to tear down the gorm db connection.
func teardownGorm(t *testing.T) {
	if gormClient != nil {
		gormClient = nil
	}
	// remove the db file
	if dbName != "" {
		err := os.Remove(dbName)
		if err != nil {
			t.Errorf("failed to remove db file: %s", err)
		}
		dbName = ""
	}
}

func TestNewGormImageRepository(t *testing.T) {
	gormClient := &gorm.DB{}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *GormImageRepository
	}{
		{
			name: "should return a new gorm image repository",
			args: args{
				db: gormClient,
			},
			want: &GormImageRepository{
				db: gormClient,
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
					gormClient = nil
				}
			}()
			t.Parallel()
			if got := NewGormImageRepository(tt.args.db); got == tt.want {
				t.Errorf("NewGormImageRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGormImageRepository_CreateImage(t *testing.T) {
	setupGorm(t)
	defer teardownGorm(t)
	type args struct {
		ctx   context.Context
		image *image.Image
	}
	tests := []struct {
		name    string
		r       *GormImageRepository
		args    args
		wantErr bool
	}{
		{
			name: "should create a new image",
			r:    NewGormImageRepository(gormClient),
			args: args{
				ctx:   context.Background(),
				image: &validImage,
			},
			wantErr: false,
		},
		{
			name: "should fail create a new image",
			r:    NewGormImageRepository(gormClient),
			args: args{
				ctx:   context.Background(),
				image: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				// recover from panic if one occured.
				if recover() != nil {
					teardownGorm(t)
				}
			}()
			if err := tt.r.CreateImage(tt.args.ctx, tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("GormImageRepository.CreateImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGormImageRepository_GetImage(t *testing.T) {
	setupGorm(t)
	defer teardownGorm(t)

	repository := NewGormImageRepository(gormClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("failed to create image: %s", err)
	}

	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		r       *GormImageRepository
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
			name: "should fail to get an image, bad account",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
			},
			want:    nil,
			wantErr: true,
			auth:    true,
		},
		{
			name: "should fail to get an image with invalid UUID",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: uuid.NewString(),
			},
			want:    nil,
			wantErr: true,
			auth:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				// recover from panic if one occured.
				if recover() != nil {
					teardownGorm(t)
				}
			}()
			config.Get().Auth = tt.auth
			got, err := tt.r.GetImage(tt.args.ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GormImageRepository.GetImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Status() != tt.want.Status() ||
				got.UUID() != tt.want.UUID() ||
				got.Name() != tt.want.Name() ||
				got.Description() != tt.want.Description() ||
				got.User() != tt.want.User() ||
				got.Distribution() != tt.want.Distribution() ||
				reflect.DeepEqual(got.OutputTypes(), tt.want.OutputTypes()) != true ||
				reflect.DeepEqual(got.Tags(), tt.want.Tags()) != true ||
				reflect.DeepEqual(got.Packages(), tt.want.Packages()) != true {
				t.Errorf("GormImageRepository.GetImage() = %v, want %v", got.Packages(), tt.want.Packages())
			}
			config.Get().Auth = false
		})
	}
}

func TestGormImageRepository_UpdateImage(t *testing.T) {
	setupGorm(t)
	defer teardownGorm(t)

	repository := NewGormImageRepository(gormClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("failed to create image: %s", err)
	}

	type args struct {
		ctx      context.Context
		uuid     string
		updateFn func(image *image.Image) (*image.Image, error)
	}
	tests := []struct {
		name    string
		r       *GormImageRepository
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
					image.SetNameAndDesc(common.Name{}, "new desc")
					return image, nil
				},
			},
			wantErr: false,
		},
		{
			name: "should fail to update an image, invalid uuid",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: uuid.NewString(),
				updateFn: func(image *image.Image) (*image.Image, error) {
					image.SetNameAndDesc(common.Name{}, "new desc")
					return image, nil
				},
			},
			wantErr: true,
		},
		{
			name: "should fail to update an image, just error",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
				updateFn: func(image *image.Image) (*image.Image, error) {
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
				t.Errorf("GormImageRepository.UpdateImage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				image, err := tt.r.GetImage(tt.args.ctx, tt.args.uuid)
				if err != nil {
					t.Errorf("failed to get image: %s", err)
				}
				if image.Description() != "new desc" {
					t.Errorf("failed to update image: %s != new desc", image.Description())
				}
			}
		})
	}
}

func TestGormImageRepository_DeleteImage(t *testing.T) {
	setupGorm(t)
	defer teardownGorm(t)

	repository := NewGormImageRepository(gormClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("failed to create image: %s", err)
	}

	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		r       *GormImageRepository
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
			auth:    false,
		},
		{
			name: "should fail to delete an image, same uuid as before",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: validImage.UUID(),
			},
			wantErr: false,
			auth:    false,
		},
		{
			name: "should fail to delete an image, invalid uuid",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: uuid.NewString(),
			},
			wantErr: false,
			auth:    false,
		},
		{
			name: "should fail to delete an image, bad account",
			r:    repository,
			args: args{
				ctx:  context.Background(),
				uuid: uuid.NewString(),
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
				t.Errorf("GormImageRepository.DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
			}
			config.Get().Auth = false
		})
	}
}

func TestGormImageRepository_GetImages(t *testing.T) {
	setupGorm(t)
	defer teardownGorm(t)

	repository := NewGormImageRepository(gormClient)
	err := repository.CreateImage(context.Background(), &validImage)
	if err != nil {
		t.Errorf("failed to create image: %s", err)
	}
	err = repository.CreateImage(context.Background(), &anotherValidImage)
	if err != nil {
		t.Errorf("failed to create image: %s", err)
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		r       *GormImageRepository
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
			auth:    false,
		},
		{
			name: "should fail to get all images, bad account",
			r:    repository,
			args: args{
				ctx: context.Background(),
			},
			want: []*image.Image{
				&validImage,
				&anotherValidImage,
			},
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
				t.Errorf("GormImageRepository.GetImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("GormImageRepository.GetImages() = %v, want %v", got, tt.want)
				}
				gotImages := make(map[string]*image.Image, len(got))
				for _, image := range got {
					gotImages[image.UUID()] = image
				}
				for _, image := range tt.want {
					if !areEqualImages(image, gotImages[image.UUID()]) {
						t.Errorf("GormImageRepository.GetImages() = %v, want %v", image, gotImages[image.UUID()])
					}
				}
			}
			config.Get().Auth = false
		})
	}
}

func TestUnmarshalPackages(t *testing.T) {
	type args struct {
		packages []models.Package
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should unmarshal packages",
			args: args{
				packages: []models.Package{
					{
						Name: "package1",
					},
					{
						Name: "package2",
					},
				},
			},
			want: []string{"package1", "package2"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := unmarshalPackages(tt.args.packages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalPackages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalTags(t *testing.T) {
	type args struct {
		tags []models.Tag
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should unmarshal tags",
			args: args{
				tags: []models.Tag{
					{
						Name: "tag1",
					},
					{
						Name: "tag2",
					},
				},
			},
			want: []string{"tag1", "tag2"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := unmarshalTags(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGormClient(t *testing.T) {
	type args struct {
		cfg *config.EdgeConfig
	}
	tests := []struct {
		name   string
		args   args
		want   *gorm.DB
		dbType string
	}{
		{
			name: "should create gorm client, mysql",
			args: args{
				cfg: config.Get(),
			},
			want:   &gorm.DB{},
			dbType: "mysql",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormClient(tt.args.cfg); got == tt.want {
				t.Errorf("NewGormClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGormAutoMigrate(t *testing.T) {
	gdb := NewGormClient(config.Get())
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
		{
			name: "should auto migrate",
			args: args{
				db: gdb,
			},
			want: gdb,
		},
		{
			name: "should panic",
			args: args{
				db: &gorm.DB{},
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
					teardownGorm(t)
				}
			}()
			if got := GormAutoMigrate(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GormAutoMigrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
