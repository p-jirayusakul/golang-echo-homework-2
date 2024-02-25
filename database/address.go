package database

import (
	"github.com/google/uuid"
)

func (x *QueriesRepository) CreateAddress(payload Address) error {

	result := x.db.Create(&payload)

	return result.Error
}

func (x *QueriesRepository) GetAddress(id uuid.UUID) (Address, error) {
	data := Address{}
	result := x.db.Where("address_id = ?", id).First(&data)
	return data, result.Error
}

func (x *QueriesRepository) UpdateAddress(arg Address) error {
	data := Address{}

	result := x.db.Where("address_id = ?", arg.AddressId).First(&data)
	if result.Error != nil {
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

func (x *QueriesRepository) DeleteAddress(id uuid.UUID) error {
	return x.db.Where("address_id = ?", id).Delete(&Address{}).Error
}
