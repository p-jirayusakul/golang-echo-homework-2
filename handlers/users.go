package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
)

func (s *ServerHttpHandler) createProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.CreateProfilesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
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

	err = s.Store.Queries.CreateProfiles(arg)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) getProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.GetProfilesByUserId)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	result, err := s.Store.Queries.GetProfiles(uid)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusCreated, result)
}

func (s *ServerHttpHandler) updateProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.UpdateProfilesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	arg := database.UpdateProfilesParams{
		UserID:    uid,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	err = s.Store.Queries.UpdateProfiles(arg)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) deleteProfiles(c echo.Context) (err error) {

	// pare json
	body := new(request.GetProfilesByUserId)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.UserID)

	err = s.Store.Queries.DeleteProfiles(uid)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) createAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.CreateAddressRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
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

	err = s.Store.Queries.CreateAddress(arg)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) getAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.GetAddressRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.AddressId)

	result, err := s.Store.Queries.GetAddress(uid)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusCreated, result)
}

func (s *ServerHttpHandler) updateAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.UpdateAddressRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
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

	err = s.Store.Queries.UpdateAddress(arg)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (s *ServerHttpHandler) deleteAddress(c echo.Context) (err error) {

	// pare json
	body := new(request.GetAddressRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	var uid uuid.UUID
	uid.Scan(body.AddressId)

	err = s.Store.Queries.DeleteAddress(uid)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}
