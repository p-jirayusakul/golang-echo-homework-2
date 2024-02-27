package response

type GetProfilesResponse struct {
	UserID    string  `json:"userId" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
	FirstName *string `json:"firstName" example:"testFirstName"`
	LastName  *string `json:"lastName" example:"testLastName"`
	Email     string  `json:"email" example:"test@email.com"`
	Phone     *string `json:"phone" example:"0888888888"`
}

type GetAddressResponse struct {
	AddressId string `json:"addressId" example:"c03217e0-6375-4bd5-bd66-b4ca5d558bc4"`
	UserID    string `json:"userId" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
	AddrType  string `json:"addrType" example:"condo"`
	AddrNo    string `json:"addrNo" example:"123/1 ABC"`
	Street    string `json:"street" example:"rachada"`
	City      string `json:"city" example:"Bangkok"`
	State     string `json:"state" example:"123456"`
}
