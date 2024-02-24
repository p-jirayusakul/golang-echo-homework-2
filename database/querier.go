package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Queries interface {
	CreateAccounts(payload Accounts) (uuid.UUID, error)
	ReadAccounts(email string) (Accounts, error)
}

type QueriesRepository struct {
	db *gorm.DB
}

func NewQueriesRepository(db *gorm.DB) QueriesRepository {
	return QueriesRepository{db: db}
}
