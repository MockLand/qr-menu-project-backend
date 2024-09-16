package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"
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

	if err := checkDishOwnership(c, userId, dishIngredients.DishId); err != nil {
		return err
	}

	if err := checkIngredientOwnership(c, userId, dishIngredients.IngredientId); err != nil {
		return err
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