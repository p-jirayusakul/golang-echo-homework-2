package configs

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Configs struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	Database string `mapstructure:"DATABASE_NAME"`
	User     string `mapstructure:"DATABASE_USER"`
	Password string `mapstructure:"DATABASE_PASSWORD"`

	HTTP_PORT  string `mapstructure:"HTTP_PORT"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
}

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func InitConfigs(filename string) *Configs {

	if _, err := os.Stat(filename); err == nil {
		viper.SetConfigFile(filename)
		viper.ReadInConfig()
	} else {
		viper.SetDefault("DATABASE_HOST", os.Getenv("DATABASE_HOST"))
		viper.SetDefault("DATABASE_PORT", ConvertInt("DATABASE_PORT"))
		viper.SetDefault("DATABASE_NAME", os.Getenv("DATABASE_NAME"))
		viper.SetDefault("DATABASE_USER", os.Getenv("DATABASE_USER"))
		viper.SetDefault("DATABASE_PASSWORD", os.Getenv("DATABASE_PASSWORD"))

		viper.SetDefault("HTTP_PORT", os.Getenv("HTTP_PORT"))
		viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))
		viper.SetDefault("SECRET_KEY", os.Getenv("SECRET_KEY"))
	}

	var config Configs
	viper.Unmarshal(&config)

	return &config
}
