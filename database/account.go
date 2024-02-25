package database

import (
	"github.com/google/uuid"
)

func (q *QueriesRepository) CreateAccounts(payload Accounts) (uuid.UUID, error) {
	result := q.db.Create(&payload)
	return payload.UserID, result.Error
}

func (x *QueriesRepository) GetAccounts(email string) (Accounts, error) {
	data := Accounts{}
	result := x.db.Where("email = ?", email).First(&data)
	return data, result.Error
}

func (x *QueriesRepository) IsEmailAlreadyExists(email string) (bool, error) {
	var count int64
	result := x.db.Model(&Accounts{}).Where("email = ?", email).Count(&count)
	return count > 0, result.Error
}
