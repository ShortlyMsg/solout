package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductView struct {
	Id uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`

	ProductID string  `gorm:"type:varchar(36);not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`

	CompanyID string  `gorm:"type:varchar(36);not null" json:"company_id"`
	Company   Company `gorm:"foreignKey:CompanyID" json:"company"`

	Price      float64    `gorm:"type:decimal(10,2);not null" json:"price"`
	ExpiryDate *time.Time `gorm:"type:date" json:"expiry_date"`

	StockQuantity int  `gorm:"not null;default:0" json:"stock_quantity"`
	SoldQuantity  int  `gorm:"not null;default:0" json:"sold_quantity"`
	IsActive      bool `gorm:"not null;default:true" json:"is_active"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
