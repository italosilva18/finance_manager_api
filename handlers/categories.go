package handlers

import (
	"finance_manager_api/database"
	"finance_manager_api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, categories)
}

func CreateCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	category.ID = uuid.New().String()
	if err := database.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.First(&category, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Category{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func GetCategoryByID(c echo.Context) error {
	id := c.Param("id")
	var category models.Category

	if err := database.DB.First(&category, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Category not found",
		})
	}

	return c.JSON(http.StatusOK, category)
}
