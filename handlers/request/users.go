package request

type CreateProfilesRequest struct {
	UserID    string `json:"userId" validate:"uuid4"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone"`
}

type GetProfilesByUserId struct {
	UserID string `param:"user_id" validate:"uuid4"`
}

type UpdateProfilesRequest struct {
	UserID    string `json:"userId" validate:"uuid4"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type CreateAddressRequest struct {
	UserID   string `json:"userId" validate:"uuid4"`
	AddrType string `json:"addrType"`
	AddrNo   string `json:"addrNo"`
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
}

type GetAddressRequest struct {
	AddressId string `param:"address_id" validate:"uuid4"`
}

type UpdateAddressRequest struct {
	AddressId string `json:"addressId" validate:"uuid4"`
	UserID    string `json:"userId" validate:"uuid4"`
	AddrType  string `json:"addrType"`
	AddrNo    string `json:"addrNo"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
}

type DeleteAddressRequest struct {
	AddressId string `param:"address_id" validate:"uuid4"`
}
