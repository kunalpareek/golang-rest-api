package config

import (
	"errors"
	"os"
)

// Config contains all application configuration
type Config struct {
	Port               string
	DbConnectionString string
}

// loadProductionConfig loads config from environment variables
func loadProductionConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		return nil, errors.New("DB credentials not found")
	}

	configuration := Config{
		Port:               port,
		DbConnectionString: dbConnectionString}

	return &configuration, nil
}

// loadDevelopmentConfig loads config from json file
func loadDevelopmentConfig() (*Config, error) {
	configuration := Config{
		Port:               "3000",
		DbConnectionString: "mongodb://localhost:27017/lookup",
	}
	return &configuration, nil
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "PRODUCTION" {
		return loadProductionConfig()
	} else {
		return loadDevelopmentConfig()
	}
}
