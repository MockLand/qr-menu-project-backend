package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"

	"github.com/labstack/echo/v4"
)

func CreateIngredient(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var ingredient model.Ingredient
	ingredient.UserId = userId

	if err := c.Bind(&ingredient); err != nil {
		c.Logger().Errorf("Failed to bind input to ingredients model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	ingredientExists := checkExistsInDB(userId, ingredient.Name, model.Ingredient{})
	if ingredientExists != nil {
		return ingredientExists
	}

	if err := database.DB.Create(&ingredient).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create ingredient"})
	}
	return c.JSON(http.StatusOK, ingredient)
}

func GetIngredients(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var ingredients []model.Ingredient
	database.DB.Where("user_id =?", userId).Find(&ingredients)

	if len(ingredients) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Ingredients not found"})
	}
	return c.JSON(http.StatusOK, ingredients)
}

func GetIngredient(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	ingredientId := c.Param("id")

	var ingredient model.Ingredient
	result := database.DB.Where("user_id =? AND ingredient_id = ?", userId, ingredientId).First(&ingredient)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve ingredient: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, ingredient)
}

func UpdateIngredient(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	ingredientId := c.Param("id")

	var ingredient model.Ingredient
	result := database.DB.Where("user_id =? AND ingredient_id =?", userId, ingredientId).First(&ingredient)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve ingredient: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&ingredient); err != nil {
		c.Logger().Errorf("Failed to bind input to ingredients model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	if err := database.DB.Save(&ingredient); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to update ingredient"})
	}
	return c.JSON(http.StatusOK, ingredient)
}

func DeleteIngredient(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	ingredientId := c.Param("id")

	var ingredient model.Ingredient
	result := database.DB.Where("user_id =? AND ingredient_id = ?", userId, ingredientId).First(&ingredient)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve ingredient: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	if err := database.DB.Delete(&ingredient).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to delete ingredient"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "Ingredient Deleted",
		"ingredient name":    ingredient.Name,
		"ingredient id":      ingredientId,
		"ingredient user id": userId,
	})
}
