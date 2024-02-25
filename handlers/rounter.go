package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/configs"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
)

type ServerHttpHandler struct {
	Cfg   *configs.Configs
	Store *database.Store
}

func NewHandler(
	app *echo.Echo,
	cfg *configs.Configs,
	store *database.Store,
) {

	handler := &ServerHttpHandler{
		Cfg:   cfg,
		Store: store,
	}

	// auth
	authGroup := app.Group("/auth")
	authGroup.POST("/register", handler.register)
	authGroup.POST("/login", handler.login)

	// users
	usersGroup := app.Group("/users")
	usersGroup.POST("/profiles", handler.createProfiles)
	usersGroup.GET("/profiles/:user_id", handler.getProfiles)
	usersGroup.PATCH("/profiles", handler.updateProfiles)
	usersGroup.DELETE("/profiles/:user_id", handler.deleteProfiles)

	usersGroup.POST("/address", handler.createAddress)
	usersGroup.GET("/address/:address_id", handler.getAddress)
	usersGroup.PATCH("/address", handler.updateAddress)
	usersGroup.DELETE("/address/:address_id", handler.deleteAddress)
}
