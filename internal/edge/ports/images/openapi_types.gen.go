// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package ports

// Defines values for Status.
const (
	StatusBuilding Status = "building"

	StatusError Status = "error"

	StatusSuccess Status = "success"
)

// CreateImageRequest defines model for CreateImageRequest.
type CreateImageRequest struct {
	Description  *Description  `json:"description,omitempty"`
	Distribution *Distribution `json:"distribution,omitempty"`
	Name         *Name         `json:"name,omitempty"`
	OutputType   *OutputTypes  `json:"output_type,omitempty"`
	Packages     *Packages     `json:"packages,omitempty"`
	Repositories *Repositories `json:"repositories,omitempty"`
	SshKey       *SSHKey       `json:"sshKey,omitempty"`
	Tags         *Tags         `json:"tags,omitempty"`
	Username     *Username     `json:"username,omitempty"`
}

// CreatedAt defines model for CreatedAt.
type CreatedAt interface{}

// DeletedAt defines model for DeletedAt.
type DeletedAt interface{}

// Description defines model for Description.
type Description string

// Distribution defines model for Distribution.
type Distribution string

// Error defines model for Error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ImageResponse defines model for ImageResponse.
type ImageResponse struct {
	CreatedAt    *CreatedAt    `json:"created_at,omitempty"`
	DeletedAt    *DeletedAt    `json:"deleted_at,omitempty"`
	Description  *Description  `json:"description,omitempty"`
	Distribution *Distribution `json:"distribution,omitempty"`
	Name         *Name         `json:"name,omitempty"`
	OutputType   *OutputTypes  `json:"output_type,omitempty"`
	Status       *Status       `json:"status,omitempty"`
	UpdatedAt    *UpdatedAt    `json:"updated_at,omitempty"`
	Uuid         *UUID         `json:"uuid,omitempty"`
	Version      *Version      `json:"version,omitempty"`
}

// Name defines model for Name.
type Name string

// OutputTypes defines model for OutputTypes.
type OutputTypes []string

// Packages defines model for Packages.
type Packages []string

// Repositories defines model for Repositories.
type Repositories []Repository

// Repository defines model for Repository.
type Repository struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// SSHKey defines model for SSHKey.
type SSHKey string

// Status defines model for Status.
type Status string

// Tags defines model for Tags.
type Tags []string

// UUID defines model for UUID.
type UUID string

// UpdateImageRequest defines model for UpdateImageRequest.
type UpdateImageRequest struct {
	Description *Description `json:"description,omitempty"`
	Name        *Name        `json:"name,omitempty"`
	Tags        *struct {
		Add    *Tags `json:"add,omitempty"`
		Remove *Tags `json:"remove,omitempty"`
	} `json:"tags,omitempty"`
}

// UpdatedAt defines model for UpdatedAt.
type UpdatedAt interface{}

// UpgradeImageRequest defines model for UpgradeImageRequest.
type UpgradeImageRequest struct {
	Description *Description `json:"description,omitempty"`
	Name        *Name        `json:"name,omitempty"`
	Packages    *struct {
		Add    *Packages `json:"add,omitempty"`
		Remove *Packages `json:"remove,omitempty"`
	} `json:"packages,omitempty"`
	Tags *struct {
		Add    *Tags `json:"add,omitempty"`
		Remove *Tags `json:"remove,omitempty"`
	} `json:"tags,omitempty"`
}

// Username defines model for Username.
type Username string

// Version defines model for Version.
type Version int

// GetImagesParams defines parameters for GetImages.
type GetImagesParams struct {
	// fields: created_at, distribution, name, status. To sort DESC use - before the fields.
	SortBy *string `json:"sort_by,omitempty"`

	// field: filter by name
	Name *string `json:"name,omitempty"`

	// field: filter by status
	Status *string `json:"status,omitempty"`
}

// CreateImageJSONBody defines parameters for CreateImage.
type CreateImageJSONBody CreateImageRequest

// UpdateImageJSONBody defines parameters for UpdateImage.
type UpdateImageJSONBody UpdateImageRequest

// CreateNewVersionJSONBody defines parameters for CreateNewVersion.
type CreateNewVersionJSONBody UpgradeImageRequest

// CreateImageJSONRequestBody defines body for CreateImage for application/json ContentType.
type CreateImageJSONRequestBody CreateImageJSONBody

// UpdateImageJSONRequestBody defines body for UpdateImage for application/json ContentType.
type UpdateImageJSONRequestBody UpdateImageJSONBody

// CreateNewVersionJSONRequestBody defines body for CreateNewVersion for application/json ContentType.
type CreateNewVersionJSONRequestBody CreateNewVersionJSONBody