package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"qr-menu-project-backend/model"
	"qr-menu-project-backend/database"
	"strings"
)

// func GetUsers(c echo.Context) error {	
// 	students, _ := GetRepoUsers()
// 	return c.JSON(http.StatusOK, students)
// }


// func GetRepoUsers() ([]model.Users, error)  {
// 	db := database.GetDBInstance()
// 	users := []model.Users{}
// 	if err:= db.Find(&users).Error; err!= nil {
// 		return nil, err
//     }
// 	return users, nil
// }

func CreateUser(c echo.Context) error {

	var user model.Users

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			return c.JSON(http.StatusConflict, map[string]interface{}{"error": "User already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, user)
}