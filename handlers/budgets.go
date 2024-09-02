package handlers

import (
	"finance_manager_api/database"
	"finance_manager_api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Criar um novo orçamento
func CreateBudget(c echo.Context) error {
	var budget models.Budget
	if err := c.Bind(&budget); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Gerar UUID para o orçamento
	budget.ID = uuid.New().String()

	if err := database.DB.Create(&budget).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, budget)
}

// Listar todos os orçamentos
func GetBudgets(c echo.Context) error {
	var budgets []models.Budget
	if err := database.DB.Find(&budgets).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, budgets)
}

// Obter um orçamento por ID
func GetBudgetByID(c echo.Context) error {
	id := c.Param("id")
	var budget models.Budget

	if err := database.DB.First(&budget, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Budget not found",
		})
	}

	return c.JSON(http.StatusOK, budget)
}

// Atualizar um orçamento por ID
func UpdateBudget(c echo.Context) error {
	id := c.Param("id")
	var budget models.Budget

	if err := database.DB.First(&budget, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Budget not found",
		})
	}

	if err := c.Bind(&budget); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Save(&budget).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, budget)
}

// Deletar um orçamento por ID
func DeleteBudget(c echo.Context) error {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Budget{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
