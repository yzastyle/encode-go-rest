package http

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host" validate:"required"`
	Port string `mapstructure:"port" validate:"required"`
}

func LoadServerConfig() (*ServerConfig, error) {
	var serverConfig ServerConfig
	if err := viper.UnmarshalKey("server", &serverConfig); err != nil {
		return nil, fmt.Errorf("error unmarshal server config: %w", err)
	}
	return &serverConfig, nil
}

func BuildServerAddress(s *ServerConfig) string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
