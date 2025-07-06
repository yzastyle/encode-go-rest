package logger

import (
	"fmt"

	"github.com/spf13/viper"
)

type LoggerConfig struct {
	ReportCaller bool `mapstructure:"report_caller"`
}

func LoadLoggerConfig() (*LoggerConfig, error) {
	var loggerConfig LoggerConfig
	if err := viper.UnmarshalKey("logging", &loggerConfig); err != nil {
		return nil, fmt.Errorf("error to unmarshal config: %w", err)
	}
	return &loggerConfig, nil
}
