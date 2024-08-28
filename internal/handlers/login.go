package handlers

import (
	"net/http"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"
	"strings"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var login_credentials model.Users

	if err := c.Bind(&login_credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}
	
	var user model.Users
	db := database.GetDBInstance()

	if err := db.Where("email  = ?", login_credentials.Email).First(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found"){
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	if login_credentials.Password!= user.Password {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid email or password"})
    }
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"user":    user,
	})
	
}
