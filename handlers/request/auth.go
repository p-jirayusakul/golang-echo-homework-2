package request

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@email.com"`
	Password string `json:"password" validate:"required,min=6" example:"123456"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@email.com"`
	Password string `json:"password" validate:"required,min=6" example:"123456"`
}
