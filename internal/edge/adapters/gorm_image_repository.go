package adapters

import (
	"context"
	"fmt"

	"github.com/Avielyo10/edge-api/internal/common/models"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/redhatinsights/edge-api/config"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GormImageRepository is a GORM implementation of the Image.Repository interface.
type GormImageRepository struct {
	db *gorm.DB
}

// NewGormImageRepository returns a new GORM implementation of the Image.Repository interface.
func NewGormImageRepository(db *gorm.DB) *GormImageRepository {
	if db == nil {
		panic("db cannot be nil")
	}
	return &GormImageRepository{db: db}
}

// CreateImage creates a new image, implementing the Image.Repository interface.
func (r *GormImageRepository) CreateImage(ctx context.Context, image *image.Image) error {
	log.Debug("gorm create image")
	return r.db.Create(image.MarshalGorm()).Error
}

// GetImage returns the image with the given UUID, implementing the Image.Repository interface.
func (r *GormImageRepository) GetImage(ctx context.Context, uuid string) (*image.Image, error) {
	log.WithField("uuid", uuid).Debug("gorm get image")
	account, err := common.GetAccountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var imageModel models.Image
	if err := r.db.Preload(clause.Associations).Where("account = ? AND uuid = ?", account.String(), uuid).First(&imageModel).Error; err != nil {
		return nil, err
	}
	newImage, err := image.UnmarshalImageFromDatabase(context.Background(),
		uuid, imageModel.Name, imageModel.Description, imageModel.Distribution, imageModel.Status,
		imageModel.User.Name, imageModel.User.SSHKey, imageModel.OutputTypes,
		unmarshalTags(imageModel.Tags), unmarshalPackages(imageModel.Packages), imageModel.Version, nil,
		imageModel.CreatedAt, imageModel.UpdatedAt, imageModel.DeletedAt.Time)
	return &newImage, err
}

// UpdateImage updates the image with the given UUID, implementing the Image.Repository interface.
func (r *GormImageRepository) UpdateImage(ctx context.Context, uuid string, updateFn func(image *image.Image) (*image.Image, error)) error {
	log.WithField("uuid", uuid).Debug("gorm update image")
	image, err := r.GetImage(ctx, uuid)
	if err != nil {
		return err
	}
	updatedImage, err := updateFn(image)
	if err != nil {
		return err
	}
	account, err := common.GetAccountFromContext(ctx)
	if err != nil {
		return err
	}
	return r.db.Model(&models.Image{}).
		Where("account = ? and uuid = ?", account.String(), image.UUID()).
		Updates(updatedImage.MarshalGorm()).Error
}

// DeleteImage deletes the image with the given UUID, implementing the Image.Repository interface.
func (r *GormImageRepository) DeleteImage(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("gorm delete image")
	account, err := common.GetAccountFromContext(ctx)
	if err != nil {
		return err
	}
	// delete image's has one/many/many2many relations when deleting an image
	return r.db.Select(clause.Associations).
		Where("account = ?", account.String()).
		Delete(&models.Image{UUID: uuid}).Error
}

// GetImages returns all images, implementing the Image.Repository interface.
func (r *GormImageRepository) GetImages(ctx context.Context) ([]*image.Image, error) {
	log.Debug("gorm get images")
	account, err := common.GetAccountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var imageModels []models.Image
	if err := r.db.Preload(clause.Associations).Where("account = ?", account.String()).Find(&imageModels).Error; err != nil {
		return nil, err
	}
	images := make([]*image.Image, len(imageModels))
	for i, imageModel := range imageModels {
		image, err := image.UnmarshalImageFromDatabase(context.Background(),
			imageModel.UUID, imageModel.Name, imageModel.Description, imageModel.Distribution, imageModel.Status,
			imageModel.User.Name, imageModel.User.SSHKey, imageModel.OutputTypes,
			unmarshalTags(imageModel.Tags), unmarshalPackages(imageModel.Packages), imageModel.Version, nil,
			imageModel.CreatedAt, imageModel.UpdatedAt, imageModel.DeletedAt.Time)
		images[i] = &image
		if err != nil {
			return nil, err
		}
	}
	return images, nil
}

// unmarshalPackages unmarshals array of package models into a string array
func unmarshalPackages(packages []models.Package) []string {
	var packagesStr []string
	for _, packageModel := range packages {
		packagesStr = append(packagesStr, packageModel.Name)
	}
	return packagesStr
}

// unmarshalTags unmarshals array of tag models into a string array
func unmarshalTags(tags []models.Tag) []string {
	var tagsStr []string
	for _, tagModel := range tags {
		tagsStr = append(tagsStr, tagModel.Name)
	}
	return tagsStr
}

// NewGormClient returns a new GormClient.
func NewGormClient(cfg *config.EdgeConfig) *gorm.DB {
	var dia gorm.Dialector
	if cfg.Database.Type == "pgsql" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
			cfg.Database.Hostname,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Name,
			cfg.Database.Port,
		)
		dia = postgres.Open(dsn)
	} else {
		dia = sqlite.Open(cfg.Database.Name)
	}
	db, err := gorm.Open(dia, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return GormAutoMigrate(db)
}

// GormAutoMigrate runs auto-migration on the given database.
func GormAutoMigrate(db *gorm.DB) *gorm.DB {
	if err := db.AutoMigrate(
		&models.Image{},
		&models.Installer{},
		&models.Tag{},
		&models.Package{},
		&models.User{},
	); err != nil {
		panic(err)
	}
	return db
}
