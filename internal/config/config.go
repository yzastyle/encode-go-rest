package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name    string `mapstructure:"name" validate:"required"`
	Version string `mapstructure:"version" validate:"required"`
}

type LoggerConfig struct {
	ReportCaller bool `mapstructure:"report_caller"`
}

func LoadConfig() (*AppConfig, *LoggerConfig, error) {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	//viper.AddConfigPath("../config")
	//viper.AddConfigPath("/config")
	viper.AddConfigPath("../../config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, nil, fmt.Errorf("error to read config file: %w", err)
	}

	var appConfig AppConfig
	if err := viper.UnmarshalKey("app", &appConfig); err != nil {
		return nil, nil, fmt.Errorf("error to unmarshal config: %w", err)
	}

	var loggerConfig LoggerConfig
	if err := viper.UnmarshalKey("logging", &loggerConfig); err != nil {
		return nil, nil, fmt.Errorf("error to unmarshal config: %w", err)
	}
	return &appConfig, &loggerConfig, nil
}
