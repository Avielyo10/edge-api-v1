package models

import (
	"time"

	"gorm.io/gorm"
)

// Model is a basic model that provides default fields for all models
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
