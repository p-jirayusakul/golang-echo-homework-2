package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers/response"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
)

// Register
// @Summary      Register By email and password
// @Description  register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body request.RegisterRequest true "body request"
// @Success      201  {object}  utils.SuccessResponse.Data{data=response.RegisterResponse}
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /auth/register [post]
func (s *ServerHttpHandler) Register(c echo.Context) (err error) {
	// pare json
	body := new(request.RegisterRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// check email before insert
	isAlreadyExists, err := s.store.IsEmailAlreadyExists(body.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if isAlreadyExists {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrEmailIsAlreadyUsed.Error())
	}

	// hash password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	arg := database.Accounts{
		Email:    body.Email,
		Password: hashedPassword,
	}

	result, err := s.store.CreateAccounts(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	argProfiles := database.Profiles{
		UserID: result,
		Email:  body.Email,
	}

	err = s.store.CreateProfiles(argProfiles)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload response.RegisterResponse
	payload.ID = result.String()
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) Login(c echo.Context) (err error) {

	// pare json
	body := new(request.LoginRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	account, err := s.store.GetAccounts(body.Email)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrLoginFail.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = utils.CheckPassword(body.Password, account.Password); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrLoginFail.Error())
	}

	var payload response.LoginResponse
	payload.Token = ""
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}
