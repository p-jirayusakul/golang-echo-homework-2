package request

type CreateProfilesRequest struct {
	UserID    string `json:"userId" validate:"uuid4" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
	FirstName string `json:"firstName" example:"testFirstName"`
	LastName  string `json:"lastName" example:"testLastName"`
	Email     string `json:"email" validate:"email" example:"test@email.com"`
	Phone     string `json:"phone" example:"0888888888"`
}

type GetProfilesByUserId struct {
	UserID string `param:"user_id" validate:"uuid4" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
}

type UpdateProfilesRequest struct {
	UserID    string `json:"userId" validate:"uuid4" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
	FirstName string `json:"firstName" validate:"required" example:"testFirstNameUpdate"`
	LastName  string `json:"lastName" validate:"required" example:"testFirstNameUpdate"`
}

type CreateAddressRequest struct {
	UserID   string `json:"userId" validate:"uuid4" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
	AddrType string `json:"addrType" example:"condo"`
	AddrNo   string `json:"addrNo" example:"123/1 ABC"`
	Street   string `json:"street" example:"rachada"`
	City     string `json:"city" example:"Bangkok"`
	State    string `json:"state" example:"123456"`
}

type GetAddressRequest struct {
	AddressId string `param:"address_id" validate:"uuid4" example:"c03217e0-6375-4bd5-bd66-b4ca5d558bc4"`
}

type UpdateAddressRequest struct {
	AddressId string `json:"addressId" validate:"uuid4" example:"c03217e0-6375-4bd5-bd66-b4ca5d558bc4"`
	UserID    string `json:"userId" validate:"uuid4" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
	AddrType  string `json:"addrType" example:"home"`
	AddrNo    string `json:"addrNo" example:"99/1 GGEZ"`
	Street    string `json:"street" example:"sukumvit"`
	City      string `json:"city" example:"Bangkok"`
	State     string `json:"state" example:"123456"`
}

type DeleteAddressRequest struct {
	AddressId string `param:"address_id" validate:"uuid4" example:"c03217e0-6375-4bd5-bd66-b4ca5d558bc4"`
}
