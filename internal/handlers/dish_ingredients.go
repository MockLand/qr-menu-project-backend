package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateDishIngredients(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var dishIngredients model.DishIngredients
	dishIngredients.UserId = userId

	if err := c.Bind(&dishIngredients); err != nil {
		c.Logger().Errorf("Failed to bind input to dish ingredients model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	if !checkDishOwnership(c, userId, dishIngredients.DishId){
		return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "User does not own the specified dish"})
	}

	if !checkIngredientOwnership(c, userId, dishIngredients.IngredientId){
		return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "User does not own the specified ingredient"})
	}

	if err := database.DB.Create(&dishIngredients).Error; err != nil {
		c.Logger().Errorf("Failed to create dish ingredients: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create dish ingredients"})
	}

	return c.JSON(http.StatusOK, dishIngredients)
}

func GetDishIngredients(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	dishId := c.Param("dish_id")

	var dishIngredients []model.DishIngredients
	if err := database.DB.Where("dish_id = ? AND user_id = ?", dishId, userId).Find(&dishIngredients).Error; err != nil {
		c.Logger().Errorf("Failed to get dish ingredients: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to get dish ingredients"})
	}

	return c.JSON(http.StatusOK, dishIngredients)
}

func DeleteDishIngredients(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	dishId := c.Param("dish_id")
	ingredientId := c.Param("ingredient_id")

	dishIdInt, err := strconv.Atoi(dishId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid dish ID"})
	}

	ingredientIdInt, err := strconv.Atoi(ingredientId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ingredient ID"})
	}

	dishIngredients := model.DishIngredients{DishId: dishIdInt, IngredientId: ingredientIdInt, UserId: userId}

	if !checkDishOwnership(c, userId, dishIngredients.DishId){
		return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "User does not own the specified dish"})
	}

	if !checkIngredientOwnership(c, userId, dishIngredients.IngredientId){
		return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "User does not own the specified ingredient"})
	}

	if err := database.DB.Where("dish_id = ? AND ingredient_id = ? AND user_id = ?", dishIngredients.DishId, dishIngredients.IngredientId, dishIngredients.UserId).First(&model.DishIngredients{}).Delete(&model.DishIngredients{}).Error; err != nil {
		c.Logger().Errorf("Failed to delete dish ingredients: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to delete dish ingredients"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Dish Ingredients Deleted",
		"dish id":       dishId,
		"ingredient id": ingredientId,
		"user id":       userId,
	})
}

