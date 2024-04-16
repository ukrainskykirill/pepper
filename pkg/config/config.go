package config

import (
	"fmt"
	"os"
	
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser string
	PostgresPassword string
	PostgresDBName string
	PostgresSSLMode string
	PostgresHost string
	PostgresPort string
}

func NewConfig() *Config{
	if err := godotenv.Load(); err != nil {
		if err := godotenv.Load("/Users/killreal/vsc_projects/pepper/tests/.env.test"); err != nil {
			fmt.Print("No .env file found")
		}
	}
	return &Config{
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDBName: os.Getenv("POSTGRES_DB_NAME"),
		PostgresSSLMode: os.Getenv("POSTGRES_SSL_MODE"),
		PostgresHost: "postgres",
		PostgresPort: os.Getenv("POSTGRES_PORT"),
	}
}