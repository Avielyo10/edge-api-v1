package query

import (
	"context"
	imageDomain "github.com/Avielyo10/edge-api/internal/edge/domain/image"
	log "github.com/sirupsen/logrus"
	"time"
)

// GetImageHandler is a handler for the Get command.
type GetImageHandler struct {
	ImageRepository imageDomain.Repository
}

// NewGetHandler returns a new GetHandler.
func NewGetImageHandler(imageRepository imageDomain.Repository) *GetImageHandler {
	if imageRepository == nil {
		return &GetImageHandler{}
	}
	return &GetImageHandler{
		ImageRepository: imageRepository,
	}
}

// Handle implements the command interface.
func (h *GetImageHandler) Handle(ctx context.Context, uuid string) (image *imageDomain.Image, err error) {
	start := time.Now()
	defer func() {
		log.
			WithError(err).
			WithField("duration", time.Since(start)).
			Debug("GetImageHandler executed")
	}()
	return h.ImageRepository.GetImage(ctx, uuid)
}
