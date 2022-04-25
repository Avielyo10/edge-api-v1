package service

import (
	"context"

	"github.com/Avielyo10/edge-api/internal/edge/adapters"
	"github.com/Avielyo10/edge-api/internal/edge/app"
	"github.com/Avielyo10/edge-api/internal/edge/app/command"
	"github.com/Avielyo10/edge-api/internal/edge/app/query"
	"github.com/redhatinsights/edge-api/config"
)

// NewApplication returns a new Application.
func NewApplication(ctx context.Context) app.Application {
	cfg := config.Get()

	redisClient := adapters.NewRedisClient(cfg)
	gormClient := adapters.NewGormClient(cfg)

	writeThroughRepository := adapters.NewReadThroughImageRepository(redisClient, gormClient)

	return app.Application{
		Commands: app.Commands{
			CreateImage:        *command.NewCreateImageHandler(writeThroughRepository),
			UpdateImage:        *command.NewUpdateImageHandler(writeThroughRepository),
			DeleteImage:        *command.NewDeleteImageHandler(writeThroughRepository),
			UpgradeImage:       *command.NewUpgradeImageHandler(writeThroughRepository),
			CancelUpgradeImage: *command.NewCancelUpgradeImageHandler(writeThroughRepository),
		},
		Queries: app.Queries{
			GetImage:  *query.NewGetImageHandler(writeThroughRepository),
			GetImages: *query.NewGetImagesHandler(writeThroughRepository),
		},
	}
}
