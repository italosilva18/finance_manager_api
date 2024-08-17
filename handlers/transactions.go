package handlers

import (
	"finance_manager_api/database"
	"finance_manager_api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	var transactions []models.Transaction
	if err := database.DB.Find(&transactions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c echo.Context) error {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	transaction.ID = uuid.New().String()
	if err := database.DB.Create(&transaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, transaction)
}

func UpdateTransaction(c echo.Context) error {
	id := c.Param("id")
	var transaction models.Transaction
	if err := database.DB.First(&transaction, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Save(&transaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Transaction{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	var transaction models.Transaction

	if err := database.DB.First(&transaction, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Transaction not found",
		})
	}

	return c.JSON(http.StatusOK, transaction)
}
