package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Queries interface {
	CreateAccounts(payload Accounts) (uuid.UUID, error)
	ReadAccounts(email string) (Accounts, error)
	CreateProfiles(payload Profiles) error
	GetProfiles(id uuid.UUID) (Profiles, error)
	UpdateProfiles(payload UpdateProfilesParams) error
	DeleteProfiles(id uuid.UUID) error
}

type QueriesRepository struct {
	db *gorm.DB
}

func NewQueriesRepository(db *gorm.DB) QueriesRepository {
	return QueriesRepository{db: db}
}
