package command

import (
	"context"

	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
)

// CancelUpgradeImageHandler is a handler for the CancelUpgradeImage command.
type CancelUpgradeImageHandler struct {
	ImageRepository image.Repository
}

// NewCancelUpgradeImageHandler returns a new CancelUpgradeImageHandler.
func NewCancelUpgradeImageHandler(imageRepository image.Repository) *CancelUpgradeImageHandler {
	if imageRepository == nil {
		return &CancelUpgradeImageHandler{}
	}
	return &CancelUpgradeImageHandler{
		ImageRepository: imageRepository,
	}
}

// Handle implements the command interface.
func (h *CancelUpgradeImageHandler) Handle(ctx context.Context, uuidToCancel string) error {
	return h.ImageRepository.UpdateImage(ctx, uuidToCancel, func(i *image.Image) (_ *image.Image, err error) {
		defer func() {
			logs.LogCommandExecution("CancelUpgradeImageHandler", uuidToCancel, err)
		}()
		if err := i.Rollback(); err != nil {
			return nil, err
		}
		return i, nil
	})
}
