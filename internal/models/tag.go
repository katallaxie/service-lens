package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tag ...
type Tag struct {
	// ID ...
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the tag name.
	Name string `json:"name" gorm:"uniqueIndex:idx_name_value" form:"name" validate:"required,min=3,max=255"`
	// Value is the tag value.
	Value string `json:"value" gorm:"uniqueIndex:idx_name_value" form:"value" validate:"required,min=3,max=255"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
