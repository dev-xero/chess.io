package config

import (
	"fmt"
	"os"
	"reflect"
)

// Config represents the server's environment variables.
type Config struct {
	ServerMode string
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPass     string
	DbName     string
	DbSslMode  string
}

// New allocates a new configuration struct in memory and
// returns a reference to the newly created struct.
// It validates the configuration fields by making sure that none
// of them are empty.
func New() (*Config, error) {
	cfg := Config{
		ServerMode: os.Getenv("SERVER_MODE"),
		Port:       os.Getenv("PORT"),
		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbPass:     os.Getenv("POSTGRES_PASS"),
		DbName:     os.Getenv("POSTGRES_NAME"),
		DbSslMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	if err := validateCfg(cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// validateCfg is a helper function that goes through each
// field of the Config struct to check for emptiness.
func validateCfg(c Config) error {
	v := reflect.ValueOf(c)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			return fmt.Errorf("'%v' is empty, have you exported your environment variables?", v.Type().Field(i).Name)
		}
	}
	return nil
}
