package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	Id uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`

	OrderID string `gorm:"type:varchar(36);not null" json:"order_id"`
	Order   Order  `gorm:"foreignKey:OrderID" json:"order"`

	ProductID string  `gorm:"type:varchar(36);not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`

	CompanyID string  `gorm:"type:varchar(36);not null" json:"company_id"`
	Company   Company `gorm:"foreignKey:CompanyID" json:"company"`

	Quantity   int     `gorm:"not null;check:quantity > 0" json:"quantity"`
	UnitPrice  float64 `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null" json:"total_price"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
