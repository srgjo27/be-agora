package config

import (
	"fmt"
	"os"
	"strconv"

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

func LoadConfig(path string) (Config, error) {
	config := Config{}

	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	config.DBHost = getEnv("DB_HOST", "localhost")
	config.DBPort = getEnv("DB_PORT", "5432")
	config.DBUser = getEnv("DB_USER", "")
	config.DBPassword = getEnv("DB_PASSWORD", "")
	config.DBName = getEnv("DB_NAME", "")
	config.DBSslMode = getEnv("DB_SSLMODE", "disable")
	config.APIPort = getEnv("API_PORT", "8080")
	config.JWTSecretKey = getEnv("JWT_SECRET_KEY", "")
	config.AccessTokenDurationMinutes = getEnvInt("JWT_ACCESS_TOKEN_DURATION_MINUTES", 60)
	config.RefreshTokenDurationHours = getEnvInt("JWT_REFRESH_TOKEN_DURATION_HOURS", 72)
	config.CookieDomain = getEnv("COOKIE_DOMAIN", "localhost")
	config.CookieSecure = getEnvBool("COOKIE_SECURE", false)

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true" || value == "1"
	}
	return defaultValue
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSslMode)
}
