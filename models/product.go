package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id        uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`
	CompanyID string    `gorm:"type:varchar(36);not null" json:"company_id"`
	Company   Company   `gorm:"foreignKey:CompanyID" json:"company"`

	CategoryID *string   `gorm:"type:varchar(36)" json:"category_id"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category"`

	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`

	ProductImages []string `gorm:"type:text[]" json:"product_images"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
