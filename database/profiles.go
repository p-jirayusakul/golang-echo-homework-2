package database

import (
	"github.com/google/uuid"
)

func (x *QueriesRepository) CreateProfiles(payload Profiles) error {
	return x.db.Create(&payload).Error
}

func (x *QueriesRepository) GetProfiles(id uuid.UUID) (Profiles, error) {
	data := Profiles{}
	result := x.db.Where("user_id = ?", id).First(&data)
	return data, result.Error
}

type UpdateProfilesParams struct {
	UserID    uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func (x *QueriesRepository) UpdateProfiles(payload UpdateProfilesParams) error {
	data := Profiles{}

	result := x.db.Where("user_id = ?", payload.UserID).First(&data)
	if result.Error != nil {
		return result.Error
	}

	data.FirstName = &payload.FirstName
	data.LastName = &payload.LastName

	return x.db.Model(Profiles{}).Where("user_id = ?", payload.UserID.String()).Updates(data).Error
}

func (x *QueriesRepository) DeleteProfiles(id uuid.UUID) error {
	return x.db.Where("user_id = ?", id.String()).Delete(&Profiles{}).Error
}
