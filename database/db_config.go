package datamanager

import "os"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

// GetConfig -> Manage the Database connection Info (This is for a postgresql connection)
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Username: os.Getenv("DB_USERNAME"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
