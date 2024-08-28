	package handlers

	import (
		"net/http"
		"qr-menu-project-backend/database"
		"qr-menu-project-backend/model"
		"strings"

		"github.com/labstack/echo/v4"
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

		createUser := database.DB.Create(&user)

		if createUser.Error != nil {
			if strings.Contains(createUser.Error.Error(), "duplicate key value violates unique constraint") {
				return c.JSON(http.StatusConflict, map[string]interface{}{"error": "User already exists"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": createUser.Error.Error()})
		}

		return c.JSON(http.StatusOK, user)
	}
