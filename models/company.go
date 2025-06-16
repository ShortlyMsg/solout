package models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	Id uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`

	UserID string `gorm:"type:varchar(36);not null;uniqueIndex" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`

	Name        string `gorm:"type:varchar(200);not null" json:"name"`
	Address     string `gorm:"type:text" json:"address"`
	Location    string `gorm:"type:varchar(200)" json:"location"`
	Phone       string `gorm:"type:varchar(20)" json:"phone"`
	Description string `gorm:"type:text" json:"description"`
	IsActive    bool   `gorm:"not null;default:true" json:"is_active"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
