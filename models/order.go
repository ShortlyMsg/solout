package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Order struct {
	Id uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`

	BuyerID string `gorm:"type:varchar(36);not null" json:"buyer_id"`
	Buyer   User   `gorm:"foreignKey:BuyerID" json:"buyer"`

	OrderNumber    string  `gorm:"type:varchar(50);not null;unique" json:"order_number"`
	TotalAmount    float64 `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	DiscountAmount float64 `gorm:"type:decimal(10,2);default:0" json:"discount_amount"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	OrderItems pq.StringArray `gorm:"foreignKey:OrderID" json:"order_items"`
	Payments   pq.StringArray `gorm:"foreignKey:OrderID" json:"payments"`
}
