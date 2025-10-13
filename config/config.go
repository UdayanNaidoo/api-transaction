package config

import (
	"log"
	"os"

	"golang.org/x/tools/go/cfg"
)

type Config struct {
	TRANSACTIONS_DB_USERNAME string 
	TRANSACTIONS_DB_PASSWORD string 
	TRANSACTIONS_DB_NAME string 
	POSTGRES_PORT string 
	POSTGRES_HOST string 
}

func LoadConfig() *Config {
	cfg = &Config{
		TRANSACTIONS_DB_USERNAME: getEnv("TRANSACTIONS_DB_USERNAME",nil),
		TRANSACTIONS_DB_PASSWORD: getEnv("TRANSACTIONS_DB_PASSWORD",nil),
		TRANSACTIONS_DB_NAME: getEnv("TRANSACTIONS_DB_NAME",nil),
		POSTGRES_PORT: getEnv("POSTGRES_PORT",nil),
		POSTGRES_HOST: getEnv("POSTGRES_HOST",nil),
	}
	return cfg
}

func getEnv(key string , defaultVal *string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}