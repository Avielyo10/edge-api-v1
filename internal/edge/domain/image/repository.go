package image

import (
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"

	"github.com/Avielyo10/edge-api/internal/common/models"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
)

var (
	// ErrImageNotFound is the error returned when the image is not found.
	ErrImageNotFound = errors.New("image not found")
)

// Repository interface for handling image data store/retrieve.
type Repository interface {
	// CreateImage creates a new image.
	CreateImage(ctx context.Context, image *Image) error
	// GetImage returns the image with the given UUID.
	GetImage(ctx context.Context, uuid string) (*Image, error)
	// UpdateImage updates the image with the given UUID.
	UpdateImage(ctx context.Context, uuid string, updateFn func(h *Image) (*Image, error)) error
	// DeleteImage deletes the image with the given UUID.
	DeleteImage(ctx context.Context, uuid string) error
	// GetImages returns all images.
	GetImages(ctx context.Context) ([]*Image, error)
}

// MarshalGorm converts a domain Image to a database Image.
func (image Image) MarshalGorm() *models.Image {
	if image.IsZero() { // if image is nil, return nil
		return nil
	}
	account, err := common.GetAccountFromContext(image.ctx)
	if err != nil {
		return nil
	}
	outputTypes := make([]string, len(image.OutputTypes()))
	for i, outputType := range image.OutputTypes() {
		outputTypes[i] = outputType.String()
	}
	model := &models.Image{
		UUID:         image.UUID(),
		Account:      account.String(),
		Name:         image.Name().String(),
		Description:  image.Description(),
		Distribution: image.Distribution().String(),
		Version:      image.Version().Uint(),
		Status:       image.Status().String(),
		OutputTypes:  outputTypes,

		Installer: *image.Installer().MarshalGorm(),
		User:      *image.User().MarshalGorm(),

		Packages: image.Packages().MarshalGorm(account.String()),
		Tags:     image.Tags().MarshalGorm(account.String()),
		Repos:    image.Repos().MarshalGorm(account.String()),
	}
	model.Installer.Account = account.String()
	model.User.Account = account.String()

	// set timing if not exists in the image
	if !image.timing.IsZero() {
		model.CreatedAt = image.timing.CreatedAt()
		model.UpdatedAt = image.timing.UpdatedAt()
		model.DeletedAt = gorm.DeletedAt{Time: image.timing.DeletedAt(), Valid: !image.timing.DeletedAt().IsZero()}

	}
	return model
}

// MarshalRedis converts a domain Image to byte array.
func (image Image) MarshalRedis() []byte {
	if image.IsZero() { // if image is nil, return nil
		return nil
	}
	imageBytes, _ := json.Marshal(image) // marshal image to byte array, marshal any error is ignored
	return imageBytes
}
