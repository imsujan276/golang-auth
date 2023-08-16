package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// AppConfig holds the application configuration.
var Config *AppConfig

type AppConfig struct {
	Debug bool   `mapstructure:"ENV"`
	Url   string `mapstructure:"APP_URL"`

	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         int    `mapstructure:"POSTGRES_PORT"`

	ServerPort int `mapstructure:"PORT"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

// LoadConfig loads the config from the given path and returns the application configuration.
func LoadConfig(configPath string) (*AppConfig, error) {
	var cfg AppConfig

	viper.AddConfigPath(configPath)
	viper.SetConfigFile("app.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %v", err)
		}
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return &cfg, nil
}