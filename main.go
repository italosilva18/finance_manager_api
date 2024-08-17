package main

import (
	"finance_manager_api/config"
	"finance_manager_api/database"
	"finance_manager_api/routes"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Aguardar alguns segundos para o banco de dados estar pronto
	time.Sleep(5 * time.Second)

	// Carregar a configuração
	config.Load()

	// Inicializar o banco de dados
	database.InitDB()

	// Criar uma nova instância do Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Definir as rotas
	routes.InitRoutes(e)

	// Iniciar o servidor
	e.Logger.Fatal(e.Start(":8080"))
}
