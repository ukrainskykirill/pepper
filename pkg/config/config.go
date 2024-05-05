package config

import (
	"fmt"
	"os"
	
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser string
	PostgresPassword string
	PostgresDB string
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
		PostgresDB: os.Getenv("POSTGRES_DB"),
		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
	}
}