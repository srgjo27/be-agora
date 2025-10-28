package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSslMode  string `mapstructure:"DB_SSLMODE"`

	APIPort string `mapstructure:"API_PORT"`

	JWTSecretKey               string `mapstructure:"JWT_SECRET_KEY"`
	AccessTokenDurationMinutes int    `mapstructure:"JWT_ACCESS_TOKEN_DURATION_MINUTES"`
	RefreshTokenDurationHours  int    `mapstructure:"JWT_REFRESH_TOKEN_DURATION_HOURS"`

	CookieDomain string `mapstructure:"COOKIE_DOMAIN"`
	CookieSecure bool   `mapstructure:"COOKIE_SECURE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSslMode)
}
