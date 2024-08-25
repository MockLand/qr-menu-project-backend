package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"qr-menu-project-backend/model"
	"qr-menu-project-backend/database"
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

	var user struct {
		ID        int    `json:"user_id" gorm:"column:user_id"`
		Username  string `json:"username" gorm:"column:username"`
		Email     string `json:"email" gorm:"column:email"`
		Password  string `json:"password_hash" gorm:"column:password_hash"`
	}

	c.Bind(&user)

	userData := model.Users{ID: user.ID, Username: user.Username, Email: user.Email, Password: user.Password}
	result := database.DB.Create(&userData)

	if result.Error!= nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }

	return c.JSON(http.StatusOK, user)
}