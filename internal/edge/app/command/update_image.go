package command

import (
	"context"

	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
)

// UpdateImage is a command to update an image.
type UpdateImage struct {
	UUIDToUpdate string
	Name         string
	Description  string
	TagsToRemove []string
	TagsToAdd    []string
}

// UpdateImageHandler is a handler for the PatchImage command.
type UpdateImageHandler struct {
	ImageRepository image.Repository
}

// NewUpdateImageHandler returns a new PatchImageHandler.
func NewUpdateImageHandler(imageRepository image.Repository) *UpdateImageHandler {
	if imageRepository == nil {
		return &UpdateImageHandler{}
	}
	return &UpdateImageHandler{
		ImageRepository: imageRepository,
	}
}

// Handle implements the command interface.
func (h *UpdateImageHandler) Handle(ctx context.Context, cmd UpdateImage) error {
	return h.ImageRepository.UpdateImage(ctx, cmd.UUIDToUpdate, func(i *image.Image) (_ *image.Image, err error) {
		defer func() {
			logs.LogCommandExecution("UpdateImageHandler", cmd, err)
		}()
		newName, err := common.NewName(cmd.Name)
		if err != nil {
			return nil, err
		}
		i.SetNameAndDesc(newName, cmd.Description)
		i.RemoveTag(common.NewTags(cmd.TagsToRemove...).Tags()...)
		i.AddTag(common.NewTags(cmd.TagsToAdd...).Tags()...)
		return i, nil
	})
}
