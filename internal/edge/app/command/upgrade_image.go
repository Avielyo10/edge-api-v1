package command

import (
	"context"

	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
)

// UpdgraeImage is a command to upgrade an image.
type UpgradeImage struct {
	UUIDToUpgrade    string
	Name             string
	Description      string
	TagsToRemove     []string
	TagsToAdd        []string
	PackagesToAdd    []string
	PackagesToRemove []string
}

// UpgradeImageHandler is a handler for the UpgradeImage command.
type UpgradeImageHandler struct {
	ImageRepository image.Repository
}

// NewUpgradeImageHandler returns a new UpgradeImageHandler.
func NewUpgradeImageHandler(imageRepository image.Repository) *UpgradeImageHandler {
	if imageRepository == nil {
		return &UpgradeImageHandler{}
	}
	return &UpgradeImageHandler{
		ImageRepository: imageRepository,
	}
}

// Handle implements the command interface.
func (h *UpgradeImageHandler) Handle(ctx context.Context, cmd UpgradeImage) error {
	return h.ImageRepository.UpdateImage(ctx, cmd.UUIDToUpgrade, func(i *image.Image) (_ *image.Image, err error) {
		defer func() {
			logs.LogCommandExecution("UpgradeImageHandler", cmd, err)
		}()
		newName, err := common.NewName(cmd.Name)
		if err != nil {
			return nil, err
		}
		i.SetNameAndDesc(newName, cmd.Description)
		i.RemoveTag(common.NewTags(cmd.TagsToRemove...).Tags()...)
		i.AddTag(common.NewTags(cmd.TagsToAdd...).Tags()...)
		i.RemovePackage(image.NewPackages(cmd.PackagesToRemove...).Packages()...)
		i.AddPackage(image.NewPackages(cmd.PackagesToAdd...).Packages()...)
		if err := i.Upgrade(); err != nil {
			return nil, err
		}
		return i, nil
	})
}
