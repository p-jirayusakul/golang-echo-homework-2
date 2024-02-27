package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
	"gorm.io/gorm"
)

func (q *Queries) CreateProfiles(payload Profiles) error {
	return q.db.Create(&payload).Error
}

func (q *Queries) GetProfiles(id uuid.UUID) (Profiles, error) {
	data := Profiles{}
	result := q.db.Where("user_id = ?", id).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Profiles{}, utils.ErrDataNotFound
		}
		return Profiles{}, result.Error
	}
	return data, result.Error
}

type UpdateProfilesParams struct {
	UserID    uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func (q *Queries) UpdateProfiles(payload UpdateProfilesParams) error {
	data := Profiles{}

	result := q.db.Where("user_id = ?", payload.UserID).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return result.Error
	}

	data.FirstName = &payload.FirstName
	data.LastName = &payload.LastName

	return q.db.Model(Profiles{}).Where("user_id = ?", payload.UserID.String()).Updates(data).Error
}

func (q *Queries) DeleteProfiles(id uuid.UUID) error {
	result := q.db.Where("user_id = ?", id.String()).Delete(&Profiles{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return result.Error
	}

	return nil
}
