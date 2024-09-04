package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateMenu(c echo.Context) error {
	// Check if the session cookie exists
	_, err := c.Cookie("session_id")
	if err != nil {
		c.Logger().Errorf("Failed to retrieve session cookie: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized request - session_id"})
	}

	// Check if user ID is available in the context
	userId, ok := UserID, true
	if !ok {
		c.Logger().Error("Failed to retrieve user_id from context")
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized request - user_id"})
	}

	// Bind the input to the menu model
	var menu model.Menus
	menu.User_id = userId
	if err := c.Bind(&menu); err != nil {
		c.Logger().Errorf("Failed to bind input to menu model: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	// Attempt to create the menu in the database
	if err := database.DB.Create(&menu).Error; err != nil {
		c.Logger().Errorf("Failed to create menu in the database: %v", err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return c.JSON(http.StatusConflict, map[string]interface{}{"error": "Menu already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create menu"})
	}

	// Return the created menu
	return c.JSON(http.StatusOK, menu)
}


func GetMenus(c echo.Context) error {
    _, err := c.Cookie("session_id")
    if err!= nil {
        c.Logger().Errorf("Failed to retrieve session cookie: %v", err)
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized request - session_id"})
    }

    userId, ok := UserID, true
    if!ok {
        c.Logger().Error("Failed to retrieve user_id from context")
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized request - user_id"})
    }

    var menus []model.Menus
    database.DB.Where("user_id =?", userId).Find(&menus)

    return c.JSON(http.StatusOK, menus)
}


