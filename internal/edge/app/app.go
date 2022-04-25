package app

import (
	"github.com/Avielyo10/edge-api/internal/edge/app/command"
	"github.com/Avielyo10/edge-api/internal/edge/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateImage        command.CreateImageHandler
	DeleteImage        command.DeleteImageHandler
	UpdateImage        command.UpdateImageHandler
	UpgradeImage       command.UpgradeImageHandler
	CancelUpgradeImage command.CancelUpgradeImageHandler
}

type Queries struct {
	GetImage  query.GetImageHandler
	GetImages query.GetImagesHandler
}
