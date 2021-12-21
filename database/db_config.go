package datamanager

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
			Host:     "localhost", // os.Getenv("DB_HOST"),
			Username: "cleverit",  // os.Getenv("DB_USERNAME"),
			Port:     "5432",      // os.Getenv("DB_PORT"),
			Password: "password",  // os.Getenv("DB_PASSWORD"),
			Name:     "cervezas",  // os.Getenv("DB_NAME"),
		},
	}
}
