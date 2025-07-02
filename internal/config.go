package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name    string `mapstructure:"name" validate:"required"`
	Version string `mapstructure:"version" validate:"required"`
}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, fmt.Errorf("error to read config file: %w", err)
	}

	var appConfig AppConfig
	if err := viper.UnmarshalKey("app", &appConfig); err != nil {
		return nil, fmt.Errorf("error to unmarshal config: %w", err)
	}

	return &appConfig, nil
}
