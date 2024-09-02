package handlers

import (
	"finance_manager_api/database"
	"finance_manager_api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Criar uma nova conta
func CreateAccount(c echo.Context) error {
	var account models.Account
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Verificar se o user_id foi fornecido
	if account.UserID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "user_id is required"})
	}

	// Gerar UUID para a conta
	account.ID = uuid.New().String()

	if err := database.DB.Create(&account).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, account)
}

// Listar todas as contas
func GetAccounts(c echo.Context) error {
	var accounts []models.Account
	if err := database.DB.Find(&accounts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, accounts)
}

// Obter uma conta por ID
func GetAccountByID(c echo.Context) error {
	id := c.Param("id")
	var account models.Account

	if err := database.DB.First(&account, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Account not found",
		})
	}

	return c.JSON(http.StatusOK, account)
}

// Atualizar uma conta por ID
func UpdateAccount(c echo.Context) error {
	id := c.Param("id")
	var account models.Account

	if err := database.DB.First(&account, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Account not found",
		})
	}

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Save(&account).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, account)
}

// Deletar uma conta por ID
func DeleteAccount(c echo.Context) error {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Account{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
