package database

import (
	"finance_manager_api/config"
	"finance_manager_api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	// Migração automática das tabelas
	DB.AutoMigrate(&models.Account{}, &models.Category{}, &models.Transaction{}, &models.User{}, &models.Budget{})
}
