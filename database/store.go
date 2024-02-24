package database

import (
	"gorm.io/gorm"
)

type Store struct {
	Queries Queries
}

func NewStore(db *gorm.DB) *Store {
	var Queries = NewQueriesRepository(db)
	return &Store{
		Queries: &Queries,
	}
}
