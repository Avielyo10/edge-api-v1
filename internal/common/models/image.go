package models

import "github.com/lib/pq"

// Image is a model for storing images.
type Image struct {
	Model

	// composite indexes (account, uuid)
	Account string `gorm:"index:idx_image,priority:1" json:"account"`
	UUID    string `gorm:"type:varchar(36);index:idx_image,priority:2" json:"uuid"`

	// image fields
	Name         string `json:"name"`
	Description  string `json:"description"`
	Distribution string `json:"distribution"`
	Status       string `json:"status"`
	Version      uint   `gorm:"default:1" json:"version"`

	User        User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Installer   Installer      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"installer"`
	Tags        []Tag          `gorm:"many2many:all_tags;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tags"`
	Packages    []Package      `gorm:"many2many:all_packages;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"packages"`
	Repos       []Repo         `gorm:"many2many:all_repos;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"repos"`
	OutputTypes pq.StringArray `gorm:"type:text[]" json:"output_types"`

	// IDs
}

// Installer is a model for storing installers.
type Installer struct {
	Model

	// index account
	Account string `gorm:"index:idx_installer,priority:1" json:"account"`

	// installer fields
	ISOURL       string `json:"iso_url"`
	ComposeJobID string `json:"compose_job_id"`
	Checksum     string `json:"checksum"`

	// IDs
	ImageID uint `json:"image_id"`
}

// Packages is a model for storing packages.
type Package struct {
	Model

	// index account
	Account string `gorm:"index:idx_packages,priority:1" json:"account"`

	// package fields
	Name string `json:"name"`

	// IDs
}

// User is a model for storing users.
type User struct {
	Model

	// index account
	Account string `gorm:"index:idx_user,priority:1" json:"account"`

	// user fields
	Name   string `json:"name"`
	SSHKey string `json:"ssh_key"`

	// IDs
	ImageID uint `json:"image_id"`
}

// Tags is a model for storing tags.
type Tag struct {
	Model

	// index account
	Account string `gorm:"index:idx_tags,priority:1" json:"account"`

	// tags fields
	Name string `json:"name"`

	// IDs
}

// Repo is a model for storing repositories.
type Repo struct {
	Model

	// index account
	Account string `gorm:"index:idx_repo,priority:1" json:"account"`

	// repository fields
	Name string `json:"name"`
	URL  string `json:"url"`

	// IDs
}
