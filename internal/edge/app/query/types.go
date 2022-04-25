package query

import "time"

// Image is a struct that represents an image.
type Image struct {
	UUID         string       `json:"uuid"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Distribution string       `json:"distribution"`
	Version      uint         `json:"version"`
	Status       string       `json:"status"`
	Installer    Installer    `json:"installer"`
	Packages     Packages     `json:"packages"`
	Repositories Repositories `json:"repositories"`
	Tags         Tags         `json:"tags"`
	User         User         `json:"user"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    time.Time    `json:"deleted_at"`
}

// Images is a list of images.
type Images struct {
	Images []Image `json:"images"`
}

// Package is a struct that represents a package.
type Package struct {
	Name string `json:"name"`
}

// Packages is a list of packages.
type Packages struct {
	Packages []Package `json:"packages"`
	UUID     string    `json:"uuid"`
}

// Installer is a struct that represents an installer.
type Installer struct {
	// TODO: Add installer fields
}

// Repository is a struct that represents a repository.
type Repository struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	UUID string `json:"uuid"`
}

// Repositories is a list of repositories.
type Repositories struct {
	Repositories []Repository `json:"repos"`
}

// Tag is a struct that represents a tag.
type Tag struct {
	Name string `json:"name"`
}

// Tags is a list of tags.
type Tags struct {
	Tags []Tag `json:"tags"`
}

// User is a struct that represents a user.
type User struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	SSHKey   string `json:"ssh_key"`
}
