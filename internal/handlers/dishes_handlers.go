package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"

	"github.com/labstack/echo/v4"
)

// eger userID categoriesteki bir userid ile eslesiyor ise eslestigi categoriese eklemeyi basariyla yap


func checkCategoryOwnership(userId int, categoryId int) bool {
    var category model.Categories
    if err := database.DB.Where("user_id = ? AND category_id = ?", userId, categoryId).First(&category).Error; err != nil {
        return false
    }
    return true
}


func CreateDish(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var dish model.Dishes
	dish.UserId = userId

	if err := c.Bind(&dish); err != nil {
		c.Logger().Errorf("Failed to bind input to dishes model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	if !checkCategoryOwnership(userId, dish.CategoryId) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "User does not own the specified category"})
    }

	dishExists := checkExistsInDB(userId, dish.Name, model.Dishes{})
	if dishExists != nil {
		return dishExists
	}

	if err := database.DB.Create(&dish).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create dish"})
	}
	return c.JSON(http.StatusOK, dish)
}

func GetDishes(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var dishes []model.Dishes
	database.DB.Where("user_id = ?", userId).Find(&dishes)

	if len(dishes) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Dishes not found"})
	}
	return c.JSON(http.StatusOK, dishes)
}

func GetDish(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	dishId := c.Param("id")

	var dish model.Dishes
	result := database.DB.Where("user_id =? AND dish_id = ?", userId, dishId).First(&dish)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve dish: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, dish)
}

func UpdateDish(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	dishId := c.Param("id")

	var dish model.Dishes
	result := database.DB.Where("user_id =? AND dish_id =?", userId, dishId).First(&dish)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve dish: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&dish); err != nil {
		c.Logger().Errorf("Failed to bind input to dishes model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	if err := database.DB.Save(&dish); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to update dish"})
	}
	return c.JSON(http.StatusOK, dish)
}

func DeleteDish(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	dishId := c.Param("id")

	var dish model.Dishes
	result := database.DB.Where("user_id =? AND dish_id =?", userId, dishId).First(&dish)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve dish: %v", result.Error)
		return c.NoContent(http.StatusNotFound)
	}

	if err := database.DB.Delete(&dish).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to delete dish"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "Dish Deleted",
		"dish name":    dish.Name,
		"dish id":      dishId,
		"dish user id": userId,
	})
}
