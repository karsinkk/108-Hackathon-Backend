package dif

import (
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

// Configuration holds all application configuration loaded from environment variables
type Configuration struct {
	APIKey         string
	FCMKey         string
	Host           string
	Port           string
	Username       string
	Password       string
	DBName         string
	VehicleFCMKey  string
	UserFCMKey     string
	ServerPort     string
	AllowedOrigins string
}

var (
	config     *Configuration
	configOnce sync.Once
)

// getEnvOrDefault retrieves an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetConfig returns the singleton configuration instance
func GetConfig() *Configuration {
	configOnce.Do(func() {
		config = &Configuration{
			APIKey:         getEnvOrDefault("GOOGLE_MAPS_API_KEY", ""),
			FCMKey:         getEnvOrDefault("FCM_KEY", ""),
			Host:           getEnvOrDefault("DB_HOST", "localhost"),
			Port:           getEnvOrDefault("DB_PORT", "5432"),
			Username:       getEnvOrDefault("DB_USERNAME", "postgres"),
			Password:       getEnvOrDefault("DB_PASSWORD", ""),
			DBName:         getEnvOrDefault("DB_NAME", "dev"),
			VehicleFCMKey:  getEnvOrDefault("VEHICLE_FCM_KEY", ""),
			UserFCMKey:     getEnvOrDefault("USER_FCM_KEY", ""),
			ServerPort:     getEnvOrDefault("SERVER_PORT", "4000"),
			AllowedOrigins: getEnvOrDefault("ALLOWED_ORIGINS", "*"),
		}

		// Log configuration (without sensitive values)
		log.Info().
			Str("db_host", config.Host).
			Str("db_port", config.Port).
			Str("db_name", config.DBName).
			Str("server_port", config.ServerPort).
			Msg("Configuration loaded")
	})
	return config
}

// ReadConf is kept for backward compatibility but uses GetConfig internally
// Deprecated: Use GetConfig() instead
func ReadConf() Configuration {
	return *GetConfig()
}
