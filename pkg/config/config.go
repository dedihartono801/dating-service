package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUser     string `envconfig:"DATABASE_USER" default:"mypostgres"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"opklnm123"`
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"postgres"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"dating"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"5432"`
}

func SetupEnvFile() *Config {
	envConfig := &Config{}
	_ = godotenv.Load()
	err := envconfig.Process("", envConfig)
	if err != nil {
		log.Fatal(nil, "Fatal error ", err)
	}

	return envConfig
}

// GetEnv
func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}
