package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
	"gorm.io/gorm"
)

func (x *Queries) CreateAddress(payload Address) error {

	result := x.db.Create(&payload)

	return result.Error
}

func (x *Queries) GetAddress(id uuid.UUID) (Address, error) {
	data := Address{}
	result := x.db.Where("address_id = ?", id).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Address{}, utils.ErrDataNotFound
		}
		return Address{}, result.Error
	}
	return data, result.Error
}

func (x *Queries) UpdateAddress(arg Address) error {
	data := Address{}

	result := x.db.Where("address_id = ?", arg.AddressId).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return result.Error
	}

	data.AddressId = arg.AddressId
	data.AddrNo = arg.AddrNo
	data.AddrType = arg.AddrType
	data.Street = arg.Street
	data.City = arg.City
	data.State = arg.State

	return x.db.Model(Address{}).Where("address_id = ?", data.AddressId).Updates(data).Error
}

func (x *Queries) DeleteAddress(id uuid.UUID) error {
	result := x.db.Where("address_id = ?", id).Delete(&Address{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return result.Error
	}
	return nil
}
