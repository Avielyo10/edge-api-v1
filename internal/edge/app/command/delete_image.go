package command

import (
	"context"
	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
)

// DeleteImageHandler is a handler for the Delete command.
type DeleteImageHandler struct {
	ImageRepository image.Repository
}

// NewDeleteImageHandler returns a new DeleteHandler.
func NewDeleteImageHandler(imageRepository image.Repository) *DeleteImageHandler {
	if imageRepository == nil {
		return &DeleteImageHandler{}
	}
	return &DeleteImageHandler{
		ImageRepository: imageRepository,
	}
}

// Handle implements the command interface.
func (h *DeleteImageHandler) Handle(ctx context.Context, uuidToDelete string) (err error) {
	defer func() {
		logs.LogCommandExecution("DeleteImageHandler", uuidToDelete, err)
	}()
	return h.ImageRepository.DeleteImage(ctx, uuidToDelete)
}
