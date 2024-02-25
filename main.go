package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-2/configs"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	app.Validator = utils.NewCustomValidator()
	app.Use(utils.ErrorHandler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogRequestID:     true,
			LogRemoteIP:      true,
			LogURI:           true,
			LogHost:          true,
			LogMethod:        true,
			LogUserAgent:     true,
			LogStatus:        true,
			LogError:         true,
			LogLatency:       true,
			LogContentLength: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				var errMsg string
				var logLevel slog.Level

				if v.Error != nil {
					errMsg = v.Error.Error()
					logLevel = slog.LevelError
				} else {
					logLevel = slog.LevelInfo
				}

				logger.LogAttrs(context.Background(), logLevel, "REQUEST",
					slog.String("id", v.RequestID),
					slog.String("remote_ip", v.RemoteIP),
					slog.String("host", v.Host),
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.String("user_agent", v.UserAgent),
					slog.Int("status", v.Status),
					slog.String("error", errMsg),
					slog.String("latency", v.Latency.String()),
				)
				return nil
			},
		}),
	)
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	handlers.NewHandler(app, cfg, store)
	app.Logger.Fatal(app.Start(cfg.HTTP_PORT))
}
