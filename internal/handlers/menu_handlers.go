package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"
	"strings"

	"github.com/labstack/echo/v4"
)

// I LITERALLY DO UNDERSTAND THIS CODE! CAN YOU EVEN IMAGINE THAT I MADE THIS CODE

func CreateMenu(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var menu model.Menus
	menu.UserId = userId
	if err := c.Bind(&menu); err != nil {
		c.Logger().Errorf("Failed to bind input to menu model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	var count int64
	database.DB.Model(&menu).Where("user_id = ? AND name = ?", userId, menu.Name).Count(&count)
	if count > 0 {
		return c.JSON(http.StatusConflict, map[string]interface{}{"error": "Menu with the same name already exists"})
	}

	if err := database.DB.Create(&menu).Error; err != nil {
		c.Logger().Errorf("Failed to create menu in the database: %v", err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return c.JSON(http.StatusConflict, map[string]interface{}{"error": "Menu already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create menu"})
	}

	return c.JSON(http.StatusOK, menu)
}

func GetMenus(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	var menus []model.Menus
	database.DB.Where("user_id =?", userId).Find(&menus)

	return c.JSON(http.StatusOK, menus)
}

func GetMenu(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	menuId := c.Param("id")

	var menu model.Menus
	result := database.DB.Where("user_id = ? AND menu_id = ?", userId, menuId).First(&menu)

	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve menu: %v", result.Error)
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Menu not found"})
	}

	return c.JSON(http.StatusOK, menu)
}

func DeleteMenu(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	menuId := c.Param("id")

	var menu model.Menus
	result := database.DB.Where("user_id =? AND menu_id =?", userId, menuId).First(&menu).Delete(&menu)

	if result.Error != nil {
		c.Logger().Errorf("Failed to delete menu: %v", result.Error)
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Menu not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "Menu Deleted",
		"menu name":    menu.Name,
		"menu id":      menuId,
		"menu user id": userId,
	})
}

func UpdateMenu(c echo.Context) error {
	userId, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	menuId := c.Param("id")

	var menu model.Menus
	result := database.DB.Where("user_id =? AND menu_id =?", userId, menuId).First(&menu)
	if result.Error != nil {
		c.Logger().Errorf("Failed to retrieve menu: %v", result.Error)
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Menu not found"})
	}

	if err := c.Bind(&menu); err != nil {
		c.Logger().Errorf("Failed to bind input to menu model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	var count int64
	database.DB.Model(&menu).Where("user_id =? AND name =?", userId, menu.Name).Count(&count)
	if count > 0 {
		return c.JSON(http.StatusConflict, map[string]interface{}{"error": "Menu with the same name already exists"})
	}

	result = database.DB.Save(&menu)
	if result.Error != nil {
		c.Logger().Errorf("Failed to update menu: %v", result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to update menu"})
	}

	return c.JSON(http.StatusOK, menu)
}

// WHO CARES THE CLEAN CODE!?!!?!?

// JK I'LL REFACTOR IT :)
