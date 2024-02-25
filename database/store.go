package database

import (
	"gorm.io/gorm"
)

type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *gorm.DB
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *gorm.DB) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
