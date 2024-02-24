package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Accounts struct {
	UserID        uuid.UUID       `gorm:"primarykey,type:uuid;default:uuid_generate_v4();uniqueIndex:pk_accounts_user_id,sort:desc"`
	ResetPassword []ResetPassword `gorm:"foreignKey:UserID;references:UserID"`
	Email         string          `gorm:"uniqueIndex:idx_accounts_email,sort:desc"`
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type ResetPassword struct {
	ResetPasswordID uuid.UUID `gorm:"primarykey,type:uuid;default:uuid_generate_v4();uniqueIndex:pk_reset_password_id,sort:desc"`
	UserID          uuid.UUID `gorm:"column:user_id"`
	RequestID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex:idx_reset_password_request_id,sort:desc"`
	ExpiresAt       time.Time
	Done            bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
