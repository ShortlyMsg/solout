package models

import (
	"time"

	"github.com/google/uuid"
)

type Favorite struct {
	Id     uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`
	UserID string    `gorm:"type:varchar(36);not null;uniqueIndex:idx_user_product_list" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`

	ProductID string  `gorm:"type:varchar(36);not null;uniqueIndex:idx_user_product_list" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`

	ListName string `gorm:"type:varchar(100);not null;default:'favorites';uniqueIndex:idx_user_product_list" json:"list_name"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
