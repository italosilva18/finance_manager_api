package handlers

import (
	"finance_manager_api/database"
	"finance_manager_api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAccounts(c echo.Context) error {
	var accounts []models.Account
	if err := database.DB.Find(&accounts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, accounts)
}

func CreateAccount(c echo.Context) error {
	var account models.Account
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	account.ID = uuid.New().String()
	if err := database.DB.Create(&account).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, account)
}

func UpdateAccount(c echo.Context) error {
	id := c.Param("id")
	var account models.Account
	if err := database.DB.First(&account, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Save(&account).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, account)
}

func DeleteAccount(c echo.Context) error {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Account{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
