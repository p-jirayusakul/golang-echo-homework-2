package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Querier interface {
	// account
	CreateAccounts(payload Accounts) (uuid.UUID, error)
	GetAccounts(email string) (Accounts, error)
	IsEmailAlreadyExists(email string) (bool, error)

	// profiles
	CreateProfiles(payload Profiles) error
	GetProfiles(id uuid.UUID) (Profiles, error)
	UpdateProfiles(payload UpdateProfilesParams) error
	DeleteProfiles(id uuid.UUID) error

	// address
	CreateAddress(payload Address) error
	GetAddress(id uuid.UUID) (Address, error)
	UpdateAddress(arg Address) error
	DeleteAddress(id uuid.UUID) error
}

type Queries struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Queries {
	return &Queries{db: db}
}

var _ Querier = (*Queries)(nil)
