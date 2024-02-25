package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers/response"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
	"gorm.io/gorm"
)

func (s *ServerHttpHandler) register(c echo.Context) (err error) {
	// pare json
	body := new(request.RegisterRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// check email before insert
	isAlreadyExists, err := s.Store.Queries.IsEmailAlreadyExists(body.Email)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	if isAlreadyExists {
		return utils.RespondWithError(c, http.StatusBadRequest, "this email is already used")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	arg := database.Accounts{
		Email:    body.Email,
		Password: hashedPassword,
	}

	result, err := s.Store.Queries.CreateAccounts(arg)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	argProfiles := database.Profiles{
		UserID: result,
		Email:  body.Email,
	}

	err = s.Store.Queries.CreateProfiles(argProfiles)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	var payload response.RegisterResponse
	payload.ID = result.String()
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) login(c echo.Context) (err error) {

	// pare json
	body := new(request.LoginRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	account, err := s.Store.Queries.GetAccounts(body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.RespondWithError(c, http.StatusUnauthorized, "email or password is invalid")
		}
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	if err = utils.CheckPassword(body.Password, account.Password); err != nil {
		return utils.RespondWithError(c, http.StatusUnauthorized, "email or password is invalid")
	}

	var payload response.LoginResponse
	payload.Token = ""
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}
