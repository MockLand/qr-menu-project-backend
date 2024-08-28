package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateMenu(c echo.Context) error {
	var menu model.Menus
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	createMenu := database.DB.Create(&menu)

	if createMenu.Error != nil {
		if strings.Contains(createMenu.Error.Error(), "duplicate key value violates unique constraint") {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create menu"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Menu already exists"})
	}
	return c.JSON(http.StatusOK, menu)

}
