package image

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
)

var (
	// ErrEmptyContext is returned when the context is empty.
	ErrEmptyContext = errors.New("empty context")
)

// Image struct represents an image of an edge device.
type Image struct {
	// context
	ctx    context.Context
	cancel context.CancelFunc
	// identity
	uuid        string
	name        common.Name
	description string
	// time
	timing common.Time
	// info
	status       Status
	version      Version
	distribution Distribution
	user         User
	// properties
	packages   Packages
	repos      Repos
	installer  Installer
	outputType []OutputType
	tags       common.Tags
}

// NewImage creates a new image.
func NewImage(uuid, name, description, distribution, status, username, sshKey string,
	outputType, tags, packages []string, version uint, repos []interface{}) (Image, error) {
	validName, err := common.NewName(name)
	if err != nil {
		return Image{}, err
	}
	validUser, err := NewUser(username, sshKey)
	if err != nil {
		return Image{}, err
	}
	validVersion, err := NewVersion(version)
	if err != nil {
		return Image{}, err
	}
	validStatus, err := NewStatusFromString(status)
	if err != nil {
		return Image{}, err
	}
	validRepos, err := ReposFromArray(repos)
	if err != nil {
		return Image{}, err
	}
	newRepos := NewRepos(validRepos...)
	image := Image{
		uuid:         uuid,
		ctx:          context.Background(),
		version:      validVersion,
		status:       validStatus,
		packages:     NewPackages(packages...),
		name:         validName,
		description:  description,
		distribution: NewDistribution(distribution),
		tags:         common.NewTags(tags...),
		user:         validUser,
		outputType:   NewOutputType(outputType...),
		repos:        newRepos,
	}
	image.WithCancel()
	return image, nil
}

// NewImageWithContext creates a new image.
func NewImageWithContext(ctx context.Context, uuid, name, description, distribution,
	status, username, sshKey string, outputType, tags, packages []string,
	version uint, repos []interface{}) (image Image, err error) {
	image, err = NewImage(uuid, name, description, distribution, status, username,
		sshKey, outputType, tags, packages, version, repos)
	if ctx == nil {
		return Image{}, ErrEmptyContext
	}
	image.ctx = ctx
	image.WithCancel()
	return
}

// IsZero returns true if the image is zero.
func (image Image) IsZero() bool {
	return image.name.IsZero() && image.description == "" &&
		image.status.IsZero() && image.version.IsZero() &&
		image.distribution.IsZero() && image.user.IsZero() &&
		reflect.DeepEqual(image.packages, Packages{}) && reflect.DeepEqual(image.repos, Repos{}) &&
		reflect.DeepEqual(image.installer, Installer{}) && len(image.outputType) == 0 &&
		reflect.DeepEqual(image.tags, common.Tags{}) && image.timing.IsZero()
}

// Account is a getter for the account of an image.
func (image Image) Account() (common.Account, error) {
	return common.GetAccountFromContext(image.ctx)
}

// UUID is a getter for the uuid of an image.
func (image Image) UUID() string {
	return image.uuid
}

// Version is a getter for a version of an image.
func (image Image) Version() Version {
	return image.version
}

// User is a getter for the user of an image.
func (image Image) User() User {
	return image.user
}

// Status is a getter for the status of an image.
func (image Image) Status() Status {
	return image.status
}

// Context is a getter for the context of an image.
func (image Image) Context() context.Context {
	return image.ctx
}

// Packages is a getter for the packages of an image.
func (image Image) Packages() Packages {
	return image.packages
}

// Repos is a getter for the repos of an image.
func (image Image) Repos() Repos {
	return image.repos
}

// Tags is a getter for the tags of an image.
func (image Image) Tags() common.Tags {
	return image.tags
}

// Name is a getter for the name of an image.
func (image Image) Name() common.Name {
	return image.name
}

// Description is a getter for the description of an image.
func (image Image) Description() string {
	return image.description
}

// Distribution is a getter for the distribution of an image.
func (image Image) Distribution() Distribution {
	return image.distribution
}

// Installer is a getter for the installer of an image.
func (image Image) Installer() Installer {
	return image.installer
}

// OutputTypes is a getter for the output type of an image.
func (image Image) OutputTypes() []OutputType {
	return image.outputType
}

// CreatedAt is a getter for the created at time of an image.
func (image Image) CreatedAt() time.Time {
	return image.timing.CreatedAt()
}

// UpdatedAt is a getter for the updated at time of an image.
func (image Image) UpdatedAt() time.Time {
	return image.timing.UpdatedAt()
}

// DeletedAt is a getter for the deleted at time of an image.
func (image Image) DeletedAt() time.Time {
	return image.timing.DeletedAt()
}

// Cancel cancels the context of an image.
func (image Image) Cancel() {
	image.cancel()
}

// SetTime sets the time of an image.
func (image *Image) SetTime(timing common.Time) {
	image.timing = timing
}

// SetNameAndDesc sets the name and description of an image.
func (image *Image) SetNameAndDesc(name common.Name, description string) {
	if !name.IsZero() {
		image.name = name
	}
	if description != "" {
		image.description = description
	}
}

// WithCancel creates a new context with a cancel function.
func (image *Image) WithCancel() {
	image.ctx, image.cancel = context.WithCancel(image.ctx)
}

// WithTimeout creates a new context with a timeout and a cancel function.
func (image *Image) WithTimeout(timeout time.Duration) {
	image.ctx, image.cancel = context.WithTimeout(image.ctx, timeout)
}

// Done returns a channel that is closed when the image update process is done.
func (image Image) Done() <-chan struct{} {
	return image.ctx.Done()
}

// MarshalJSON creates a custom JSON marshaler for an image.
func (image Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		UUID         string       `json:"uuid,omitempty"`
		Name         common.Name  `json:"name,omitempty"`
		Description  string       `json:"description,omitempty"`
		Status       Status       `json:"status,omitempty"`
		Version      Version      `json:"version,omitempty"`
		Distribution Distribution `json:"distribution,omitempty"`
		User         User         `json:"user,omitempty"`
		Packages     Packages     `json:"packages,omitempty"`
		Repos        Repos        `json:"repos,omitempty"`
		Installer    Installer    `json:"installer,omitempty"`
		OutputType   []OutputType `json:"outputType,omitempty"`
		Tags         common.Tags  `json:"tags,omitempty"`
		CreatedAt    string       `json:"created_at,omitempty"`
		UpdatedAt    string       `json:"updated_at,omitempty"`
		DeletedAt    string       `json:"deleted_at,omitempty"`
	}{
		UUID:         image.uuid,
		Name:         image.name,
		Description:  image.description,
		Status:       image.status,
		Version:      image.version,
		Distribution: image.distribution,
		User:         image.user,
		Packages:     image.packages,
		Repos:        image.repos,
		Installer:    image.installer,
		OutputType:   image.outputType,
		Tags:         image.tags,
		CreatedAt:    image.timing.CreatedAt().Format(time.RFC3339Nano),
		UpdatedAt:    image.timing.UpdatedAt().Format(time.RFC3339Nano),
		DeletedAt:    image.timing.DeletedAt().Format(time.RFC3339Nano),
	})
}

// UnmarshalJSON unmarshals the image from JSON
func (image *Image) UnmarshalJSON(data []byte) error {
	var imageData struct {
		UUID         string       `json:"uuid,omitempty"`
		Name         common.Name  `json:"name,omitempty"`
		Description  string       `json:"description,omitempty"`
		Status       Status       `json:"status,omitempty"`
		Version      Version      `json:"version,omitempty"`
		Distribution Distribution `json:"distribution,omitempty"`
		User         User         `json:"user,omitempty"`
		Packages     Packages     `json:"packages,omitempty"`
		Repos        Repos        `json:"repos,omitempty"`
		Installer    Installer    `json:"installer,omitempty"`
		OutputType   []OutputType `json:"outputType,omitempty"`
		Tags         common.Tags  `json:"tags,omitempty"`
		CreatedAt    string       `json:"created_at,omitempty"`
		UpdatedAt    string       `json:"updated_at,omitempty"`
		DeletedAt    string       `json:"deleted_at,omitempty"`
	}
	err := json.Unmarshal(data, &imageData)
	if err != nil {
		return err
	}
	image.uuid = imageData.UUID
	image.name = imageData.Name
	image.description = imageData.Description
	image.status = imageData.Status
	image.version = imageData.Version
	image.distribution = imageData.Distribution
	image.user = imageData.User
	image.packages = imageData.Packages
	image.repos = imageData.Repos
	image.installer = imageData.Installer
	image.outputType = imageData.OutputType
	image.tags = imageData.Tags

	createdAt, err := time.Parse(time.RFC3339Nano, imageData.CreatedAt)
	if err != nil {
		return err
	}
	updatedAt, err := time.Parse(time.RFC3339Nano, imageData.UpdatedAt)
	if err != nil {
		return err
	}
	deletedAt, err := time.Parse(time.RFC3339Nano, imageData.DeletedAt)
	if err != nil {
		return err
	}
	image.timing = common.NewTime(createdAt, updatedAt, deletedAt)
	return nil
}

// UnmarshalImageFromDatabase unmarshals the image from the database.
func UnmarshalImageFromDatabase(ctx context.Context, uuid, name, description, distribution,
	status, username, sshKey string,
	outputType, tags, packages []string, version uint, repos []interface{},
	createdAt, updatedAt, deletedAt time.Time) (image Image, err error) {
	image, err = NewImageWithContext(ctx, uuid, name, description, distribution,
		status, username, sshKey, outputType, tags, packages, version, repos)
	image.SetTime(common.NewTime(createdAt, updatedAt, deletedAt))
	return
}
