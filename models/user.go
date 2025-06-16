package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `gorm:"type:varchar(36);default:uuid_generate_v4();primaryKey" json:"id"`
	Email     string     `gorm:"type:varchar(80);not null;uniqueIndex" json:"email"`
	Password  string     `gorm:"type:varchar(32);not null" json:"-"` // "-" ile JSON'da gizleyebiliyoz
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Phone     *string    `gorm:"type:varchar(20)" json:"phone"`
	BirthDate *time.Time `gorm:"type:date" json:"birth_date"`
	Gender    string     `gorm:"type:varchar(10);check:gender IN ('male','female','other')" json:"gender"`
	UserType  string     `gorm:"type:varchar(20);not null;default:'customer';check:user_type IN ('customer','company','admin')" json:"user_type"`

	IsActive             bool       `gorm:"not null;default:true" json:"is_active"`
	EmailVerified        bool       `gorm:"not null;default:false" json:"email_verified"`
	PasswordResetToken   *string    `gorm:"type:varchar(255)" json:"-"`
	PasswordResetExpires *time.Time `gorm:"type:timestamp" json:"-"`
	LastLoginAt          *time.Time `gorm:"type:timestamp" json:"last_login_at"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
