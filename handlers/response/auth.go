package response

type RegisterResponse struct {
	ID string `json:"id" example:"6baad1bd-ef73-4092-a7ae-afb72f9921e3"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
