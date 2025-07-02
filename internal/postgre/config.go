package postgre

import (
	"fmt"

	"github.com/spf13/viper"
)

type DataSourceConfig struct {
	Type     string `mapstructure:"type" validate:"required"`
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DbName   string `mapstructure:"dbname" validate:"required"`
	SSLMode  string `mapstructure:"ssl_mode" validate:"required"`
}

func LoadDataSourceConfig() (*DataSourceConfig, error) {
	var dataSourceConfig DataSourceConfig
	if err := viper.UnmarshalKey("datasource", &dataSourceConfig); err != nil {
		return nil, fmt.Errorf("error unmarshal datasource config: %w", err)
	}
	return &dataSourceConfig, nil
}

func BuildConnectionURL(d *DataSourceConfig) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", d.Type, d.User, d.Password, d.Host, d.Port, d.DbName, d.SSLMode)
}
