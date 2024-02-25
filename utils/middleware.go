package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator holds the validator instance.
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new instance of CustomValidator.
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validates the given struct using the validator instance.
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			// Handle errors here
			switch e := err.(type) {
			case *echo.HTTPError:
				return RespondWithError(c, e.Code, e.Message.(string))
			default:
				return RespondWithError(c, http.StatusInternalServerError, "Internal Server Error")
			}
		}
		return nil
	}
}
