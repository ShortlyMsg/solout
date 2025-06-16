package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id          uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`

	ParentID *string   `gorm:"type:varchar(36)" json:"parent_id"`
	Parent   *Category `gorm:"foreignKey:ParentID;" json:"parent"`
	IsActive bool      `gorm:"not null;default:true" json:"is_active"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
