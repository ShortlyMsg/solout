package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id uuid.UUID `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`

	OrderID string `gorm:"type:varchar(36);not null" json:"order_id"`
	Order   Order  `gorm:"foreignKey:OrderID" json:"order"`

	Amount          float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod   string  `gorm:"type:varchar(30)" json:"payment_method"`
	PaymentStatus   string  `gorm:"type:varchar(20)" json:"payment_status"`
	PaymentResponse string  `gorm:"type:text" json:"payment_response"`
	TransactionID   string  `gorm:"type:varchar(100)" json:"transaction_id"`

	ProcessedAt *time.Time `gorm:"type:timestamp" json:"processed_at"`
	CreatedAt   time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
