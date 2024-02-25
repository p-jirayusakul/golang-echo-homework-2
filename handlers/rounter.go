package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/configs"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
)

type ServerHttpHandler struct {
	cfg   *configs.Configs
	store database.Store
}

func NewHandler(
	app *echo.Echo,
	cfg *configs.Configs,
	store database.Store,
) *ServerHttpHandler {

	handler := &ServerHttpHandler{
		cfg:   cfg,
		store: store,
	}

	// auth
	authGroup := app.Group("/auth")
	authGroup.POST("/register", handler.Register)
	authGroup.POST("/login", handler.Login)

	// users
	usersGroup := app.Group("/users")
	usersGroup.POST("/profiles", handler.CreateProfiles)
	usersGroup.GET("/profiles/:user_id", handler.GetProfiles)
	usersGroup.PATCH("/profiles", handler.UpdateProfiles)
	usersGroup.DELETE("/profiles/:user_id", handler.DeleteProfiles)

	usersGroup.POST("/address", handler.CreateAddress)
	usersGroup.GET("/address/:address_id", handler.GetAddress)
	usersGroup.PATCH("/address", handler.UpdateAddress)
	usersGroup.DELETE("/address/:address_id", handler.DeleteAddress)

	return handler
}
