package config

import (
	"os"
)

type Config struct {
	TRANSACTIONS_DB_USERNAME string 
	TRANSACTIONS_DB_PASSWORD string 
	TRANSACTIONS_DB_NAME string 
	POSTGRES_PORT string 
	POSTGRES_HOST string 
}

func LoadConfig() *Config {
	cfg := &Config{
		TRANSACTIONS_DB_USERNAME: getEnv("TRANSACTIONS_DB_USERNAME"),
		TRANSACTIONS_DB_PASSWORD: getEnv("TRANSACTIONS_DB_PASSWORD"),
		TRANSACTIONS_DB_NAME: getEnv("TRANSACTIONS_DB_NAME"),
		POSTGRES_PORT: getEnv("POSTGRES_PORT"),
		POSTGRES_HOST: getEnv("POSTGRES_HOST"),
	}
	return cfg
} 
func getEnv(key string , defaultVal ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
		if len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return ""
}