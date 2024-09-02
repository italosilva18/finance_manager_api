package database

import (
	"finance_manager_api/config"
	"finance_manager_api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize initializes the database connection and performs migrations
func Initialize() {
	var err error

	// Conecte-se ao banco de dados usando o GORM e o driver PostgreSQL
	DB, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	// Realize as migrações automáticas para criar as tabelas, se necessário
	err = DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Category{}, &models.Transaction{}, &models.Budget{})
	if err != nil {
		log.Fatalf("Falha ao migrar banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida e migrações aplicadas")
}
