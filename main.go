package main

import (
	"finance_manager_api/config"
	"finance_manager_api/database"
	"finance_manager_api/routes"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Carregar configurações
	config.Load()

	// Inicialize o banco de dados
	database.Initialize()

	e := echo.New()

	// Middleware CORS (já existente)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:49837", "http://localhost:8080"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Middleware JWT para proteger as rotas
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//     SigningKey:  []byte("dX6N2!mc5pZ9@k1Y3QwVr7HbTx$zA8J0"), // Substitua pela sua chave secreta
	//     TokenLookup: "header:Authorization",
	//     AuthScheme:  "Bearer",
	// }))

	// Rotas
	routes.InitRoutes(e)

	// Inicie o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
