package routes

import (
	"finance_manager_api/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Rotas para transações
	e.GET("/transactions", handlers.GetTransactions)
	e.POST("/transactions", handlers.CreateTransaction)
	e.PUT("/transactions/:id", handlers.UpdateTransaction)
	e.DELETE("/transactions/:id", handlers.DeleteTransaction)
	e.GET("/transactions/:id", handlers.GetTransactionByID)

	// Rotas para orçamentos
	e.GET("/budgets", handlers.GetBudgets)
	e.POST("/budgets", handlers.CreateBudget)
	e.PUT("/budgets/:id", handlers.UpdateBudget)
	e.DELETE("/budgets/:id", handlers.DeleteBudget)

	// Rotas para usuários
	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	// Rotas para categorias
	e.GET("/categories", handlers.GetCategories)
	e.POST("/categories", handlers.CreateCategory)
	e.PUT("/categories/:id", handlers.UpdateCategory)
	e.DELETE("/categories/:id", handlers.DeleteCategory)

	// Rotas para contas
	e.GET("/accounts", handlers.GetAccounts)
	e.POST("/accounts", handlers.CreateAccount)
	e.PUT("/accounts/:id", handlers.UpdateAccount)
	e.DELETE("/accounts/:id", handlers.DeleteAccount)

	// Rotas para budgets
	e.GET("/budgets", handlers.GetBudgets)
	e.POST("/budgets", handlers.CreateBudget)
	e.PUT("/budgets/:id", handlers.UpdateBudget)
	e.DELETE("/budgets/:id", handlers.DeleteBudget)

}
