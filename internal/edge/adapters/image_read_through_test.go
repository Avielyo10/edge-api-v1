package adapters

import (
	"context"
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

var (
	validImage, _ = image.NewImageWithContext(
		context.Background(),
		uuid.NewString(),
		"test-image",
		"some description",
		"rhel8",
		"success",
		"redhat-user",
		"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey",
		[]string{"rhel-edge-commit"},
		[]string{"tag1", "tag2"},
		[]string{"vim", "emacs"},
		1,
		[]interface{}{
			map[string]interface{}{
				"name": "test",
				"url":  "http://test.com",
			},
			map[string]interface{}{
				"name": "test2",
				"url":  "http://test2.com",
			},
		},
	)
	anotherValidImage, _ = image.NewImageWithContext(
		context.Background(),
		uuid.NewString(),
		"test-image-2",
		"some description 2",
		"rhel8",
		"success",
		"redhat-user",
		"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey",
		[]string{"rhel-edge-commit"},
		[]string{"tag1", "tag2"},
		[]string{"vim", "emacs"},
		1,
		[]interface{}{
			map[string]interface{}{
				"name": "test",
				"url":  "http://test.com",
			},
			map[string]interface{}{
				"name": "test2",
				"url":  "http://test2.com",
			},
		},
	)
)

func TestNewReadThroughImageRepository(t *testing.T) {
	type args struct {
		rdb *redis.Client
		gdb *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *ReadThroughImageRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewReadThroughImageRepository(tt.args.rdb, tt.args.gdb); got == tt.want {
				t.Errorf("NewReadThroughImageRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadThroughImageRepository_CreateImage(t *testing.T) {
	type args struct {
		ctx   context.Context
		image *image.Image
	}
	tests := []struct {
		name    string
		r       *ReadThroughImageRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.r.CreateImage(tt.args.ctx, tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("ReadThroughImageRepository.CreateImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadThroughImageRepository_GetImage(t *testing.T) {
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		r       *ReadThroughImageRepository
		args    args
		want    *image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.r.GetImage(tt.args.ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadThroughImageRepository.GetImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadThroughImageRepository.GetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadThroughImageRepository_UpdateImage(t *testing.T) {
	type args struct {
		ctx      context.Context
		uuid     string
		updateFn func(image *image.Image) (*image.Image, error)
	}
	tests := []struct {
		name    string
		r       *ReadThroughImageRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.r.UpdateImage(tt.args.ctx, tt.args.uuid, tt.args.updateFn); (err != nil) != tt.wantErr {
				t.Errorf("ReadThroughImageRepository.UpdateImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadThroughImageRepository_DeleteImage(t *testing.T) {
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		r       *ReadThroughImageRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.r.DeleteImage(tt.args.ctx, tt.args.uuid); (err != nil) != tt.wantErr {
				t.Errorf("ReadThroughImageRepository.DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadThroughImageRepository_GetImages(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		r       *ReadThroughImageRepository
		args    args
		want    []*image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.r.GetImages(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadThroughImageRepository.GetImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadThroughImageRepository.GetImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

// utillity function to compare two images, return true if they are equal
func areEqualImages(a, b *image.Image) bool {
	if a.Status() != b.Status() ||
		a.UUID() != b.UUID() ||
		a.Name() != b.Name() ||
		a.Description() != b.Description() ||
		a.User() != b.User() ||
		a.Distribution() != b.Distribution() ||
		reflect.DeepEqual(a.OutputTypes(), b.OutputTypes()) != true ||
		reflect.DeepEqual(a.Tags(), b.Tags()) != true ||
		reflect.DeepEqual(a.Packages(), b.Packages()) != true {
		return false
	}
	return true
}
