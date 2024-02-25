package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
)

func (s *ServerHttpHandler) CreateProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.CreateProfilesRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	arg := database.Profiles{
		UserID:    uid,
		FirstName: &body.FirstName,
		LastName:  &body.LastName,
		Email:     body.Email,
		Phone:     &body.Phone,
	}

	err = s.store.CreateProfiles(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) GetProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.GetProfilesByUserId)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	result, err := s.store.GetProfiles(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusCreated, result)
}

func (s *ServerHttpHandler) UpdateProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.UpdateProfilesRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	arg := database.UpdateProfilesParams{
		UserID:    uid,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	err = s.store.UpdateProfiles(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (s *ServerHttpHandler) DeleteProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.GetProfilesByUserId)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	err = s.store.DeleteProfiles(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) CreateAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.CreateAddressRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	arg := database.Address{
		UserID:   uid,
		AddrType: body.AddrType,
		AddrNo:   body.AddrNo,
		Street:   body.Street,
		City:     body.City,
		State:    body.State,
	}

	err = s.store.CreateAddress(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) GetAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.GetAddressRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.AddressId)

	result, err := s.store.GetAddress(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusCreated, result)
}

func (s *ServerHttpHandler) UpdateAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.UpdateAddressRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var aid uuid.UUID
	aid.Scan(body.AddressId)

	var uid uuid.UUID
	uid.Scan(body.UserID)

	arg := database.Address{
		AddressId: aid,
		UserID:    uid,
		AddrType:  body.AddrType,
		AddrNo:    body.AddrNo,
		Street:    body.Street,
		City:      body.City,
		State:     body.State,
	}

	err = s.store.UpdateAddress(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) DeleteAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.GetAddressRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.AddressId)

	err = s.store.DeleteAddress(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}
