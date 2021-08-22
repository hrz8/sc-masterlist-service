package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	SERVICE_PORT int    `mapstructure:"SERVICE_PORT"`
	DB_HOST      string `mapstructure:"DB_HOST"`
	DB_PORT      int    `mapstructure:"DB_PORT"`
	DB_USER      string `mapstructure:"DB_USER"`
	DB_PASSWORD  string `mapstructure:"DB_PASSWORD"`
	DB_NAME      string `mapstructure:"DB_NAME"`
}

var appConfig AppConfig

// NewConfig return configurations implementation
func NewConfig() (*AppConfig, error) {
	v := viper.New()

	v.AddConfigPath("..")
	v.SetConfigName("config")
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	v.Unmarshal(&appConfig)
	return &appConfig, nil
}
