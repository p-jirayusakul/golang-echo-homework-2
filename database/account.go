package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
	"gorm.io/gorm"
)

func (q *Queries) CreateAccounts(payload Accounts) (uuid.UUID, error) {
	result := q.db.Create(&payload)
	return payload.UserID, result.Error
}

func (q *Queries) GetAccounts(email string) (Accounts, error) {
	data := Accounts{}
	result := q.db.Where("email = ?", email).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Accounts{}, utils.ErrDataNotFound
		}
		return Accounts{}, result.Error
	}
	return data, result.Error
}

func (q *Queries) IsEmailAlreadyExists(email string) (bool, error) {
	var count int64
	result := q.db.Model(&Accounts{}).Where("email = ?", email).Count(&count)
	return count > 0, result.Error
}
