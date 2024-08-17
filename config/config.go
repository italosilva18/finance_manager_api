package config

import (
	"fmt"
	"os"
)

var (
	DBHost     = os.Getenv("DB_HOST")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
	DSN        = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBName, DBPassword)
)

func Load() {
	// Carregar outras configurações, se necessário
}
