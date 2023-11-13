package adapters

import (
	"gorm.io/gorm"

	"github.com/katallaxie/service-lense/internal/ports"
)

var _ ports.LensRepository = (*DB)(nil)

// DB ...
type DB struct {
	db *gorm.DB
	ports.LensRepositoryUnimplemented
}

// NewDB ...
func NewDB(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}