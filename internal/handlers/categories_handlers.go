package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"

	"github.com/labstack/echo/v4"
)

func CreateCategory(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var category model.Categories
	category.UserId = userId

	if err := c.Bind(&category); err != nil {
		c.Logger().Errorf("Failed to bind input to categories model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	categoryExists := checkExistsInDB(userId, category.Name, model.Categories{})
	if categoryExists != nil {
		return categoryExists
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Faileddd to create category"})
	}
	return c.JSON(http.StatusOK, category)
}

func GetCategories(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var categories []model.Categories
	database.DB.Where("user_id =?", userId).Find(&categories)

	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	categoryId := c.Param("id")

	var category model.Categories
	result := database.DB.Where("user_id =? AND category_id = ?", userId, categoryId).First(&category)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve category: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	categoryId := c.Param("id")

	var category model.Categories
	result := database.DB.Where("user_id =? AND category_id = ?", userId, categoryId).First(&category)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve category: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	categoryExists := checkExistsInDB(userId, category.Name, model.Categories{})
	if categoryExists != nil {
		return categoryExists
	}

	if err := database.DB.Save(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to update category"})
	}
	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	categoryId := c.Param("id")

	var category model.Categories
	result := database.DB.Where("user_id =? AND category_id = ?", userId, categoryId).First(&category)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve category: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to delete category"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "Menu Deleted",
		"menu name":    category.Name,
		"menu id":      categoryId,
		"menu user id": userId,
	})
}
