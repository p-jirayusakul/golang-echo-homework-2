package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/configs"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			// Handle errors here
			switch e := err.(type) {
			case *echo.HTTPError:
				return utils.RespondWithError(c, e.Code, e.Message.(string))
			default:
				return utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error")
			}
		}
		return nil
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {

	cfg := configs.InitConfigs(".env")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&database.Accounts{}, &database.ResetPassword{}, &database.Address{}, &database.Profiles{})

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	store := database.NewStore(db)

	app := echo.New()
	app.Validator = &CustomValidator{validator: validator.New()}
	app.Use(ErrorHandler)
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	handlers.NewHandler(app, cfg, store)
	app.Logger.Fatal(app.Start(cfg.HTTP_PORT))
}
