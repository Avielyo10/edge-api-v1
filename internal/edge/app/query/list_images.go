package query

import (
	"context"
	imageDomain "github.com/Avielyo10/edge-api/internal/edge/domain/image"
	log "github.com/sirupsen/logrus"
	"time"
)

// GetImagesHandler is a handler for the Get command.
type GetImagesHandler struct {
	ImageRepository imageDomain.Repository
}

// NewGetHandler returns a new GetHandler.
func NewGetImagesHandler(imageRepository imageDomain.Repository) *GetImagesHandler {
	if imageRepository == nil {
		return &GetImagesHandler{}
	}
	return &GetImagesHandler{
		ImageRepository: imageRepository,
	}
}

// Handle implements the command interface.
func (h *GetImagesHandler) Handle(ctx context.Context) (images []*imageDomain.Image, err error) {
	start := time.Now()
	defer func() {
		log.
			WithError(err).
			WithField("duration", time.Since(start)).
			Debug("GetImagesHandler executed")
	}()
	return h.ImageRepository.GetImages(ctx)
}
