package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// service core config
	RestPort string `envconfig:"REST_PORT" default:"8080"`

	// database config
	DBUsername string `envconfig:"DB_USERNAME" default:"infokes_user"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"infokes_password"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"3000"`
	DBName     string `envconfig:"DB_NAME" default:"infokes_db"`
}

// Get to get defined configuration
func Get() Config {
	godotenv.Load()
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
