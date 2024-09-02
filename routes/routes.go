package routes

import (
	"finance_manager_api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	// Rota de login (não protegida)
	e.POST("/users/login", handlers.LoginUser)

	// Grupo de rotas protegidas por JWT
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("dX6N2!mc5pZ9@k1Y3QwVr7HbTx$zA8J0"),
	}))

	// Usuários
	api.GET("/users", handlers.GetUsers)
	api.POST("/users", handlers.CreateUser)
	api.GET("/users/:id", handlers.GetUserByID)
	api.PUT("/users/:id", handlers.UpdateUser)
	api.DELETE("/users/:id", handlers.DeleteUser)

	// Contas
	api.POST("/accounts", handlers.CreateAccount)
	api.GET("/accounts", handlers.GetAccounts)
	api.GET("/accounts/:id", handlers.GetAccountByID)
	api.PUT("/accounts/:id", handlers.UpdateAccount)
	api.DELETE("/accounts/:id", handlers.DeleteAccount)

	// Categorias
	api.POST("/categories", handlers.CreateCategory)
	api.GET("/categories", handlers.GetCategories)
	api.GET("/categories/:id", handlers.GetCategoryByID)
	api.PUT("/categories/:id", handlers.UpdateCategory)
	api.DELETE("/categories/:id", handlers.DeleteCategory)

	// Transações
	api.POST("/transactions", handlers.CreateTransaction)
	api.GET("/transactions", handlers.GetTransactions)
	api.GET("/transactions/:id", handlers.GetTransactionByID)
	api.PUT("/transactions/:id", handlers.UpdateTransaction)
	api.DELETE("/transactions/:id", handlers.DeleteTransaction)

	// Orçamentos
	api.POST("/budgets", handlers.CreateBudget)
	api.GET("/budgets", handlers.GetBudgets)
	api.GET("/budgets/:id", handlers.GetBudgetByID)
	api.PUT("/budgets/:id", handlers.UpdateBudget)
	api.DELETE("/budgets/:id", handlers.DeleteBudget)

}
