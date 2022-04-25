package command

import (
	"context"
	"time"

	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
)

// CreateImage is a command to create an image.
type CreateImage struct {
	UUID         string
	Status       string
	Name         string
	Description  string
	Distribution string
	Username     string
	SSHKey       string
	OutputType   []string
	Tags         []string
	Packages     []string
	Version      uint
	Repos        []interface{}
}

// CreateImageHandler is a handler for the CreateImage command.
type CreateImageHandler struct {
	ImageRepository image.Repository
}

// NewCreateImageHandler returns a new CreateImageHandler.
func NewCreateImageHandler(imageRepository image.Repository) *CreateImageHandler {
	if imageRepository == nil {
		return &CreateImageHandler{}
	}
	return &CreateImageHandler{
		ImageRepository: imageRepository,
	}
}

// Handle implements the command interface.
func (h *CreateImageHandler) Handle(ctx context.Context, cmd CreateImage) (_ *image.Image, err error) {
	defer func() {
		logs.LogCommandExecution("CreateImageHandler", cmd, err)
	}()
	newImage, err := image.NewImageWithContext(ctx, cmd.UUID, cmd.Name,
		cmd.Description, cmd.Distribution, cmd.Status, cmd.Username, cmd.SSHKey, cmd.OutputType, cmd.Tags, cmd.Packages, cmd.Version, cmd.Repos)
	if err != nil {
		return nil, err
	}
	newImage.SetTime(common.NewTime(time.Now(), time.Now(), time.Time{}))
	// TODO: image-builder should create the image here
	return &newImage, h.ImageRepository.CreateImage(ctx, &newImage)
}
