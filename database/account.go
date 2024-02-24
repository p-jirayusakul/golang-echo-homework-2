package database

import (
	"github.com/google/uuid"
)

func (q *QueriesRepository) CreateAccounts(payload Accounts) (uuid.UUID, error) {
	result := q.db.Create(&payload)
	return payload.UserID, result.Error
}

func (x *QueriesRepository) ReadAccounts(email string) (Accounts, error) {
	data := Accounts{}
	result := x.db.Where("email = ?", email).First(&data)
	return data, result.Error
}
