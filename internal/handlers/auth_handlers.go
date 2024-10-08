package handlers

import (
	"net/http"
	"os"
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var UserID int

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func getSessionAndUserID(c echo.Context) (int, error) {
	_, err := c.Cookie("session_id")
	if err != nil {
		c.Logger().Errorf("Failed to retrieve session cookie: %v", err)
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized request - session_id")
	}

	userId, ok := UserID, true
	if !ok {
		c.Logger().Error("Failed to retrieve user_id from context")
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized request - user_id")
	}

	return userId, nil
}

func Login(c echo.Context) error {
	c.Logger().Info("Login function started")
	var loginCredentials model.LoginCredentials

	if err := c.Bind(&loginCredentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	foundUser := model.Users{}
	result := database.DB.Where("email = ?", loginCredentials.Email).First(&foundUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "No user found"})
	}

	if foundUser.Password != loginCredentials.Password {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid email or password"})
	}

	expirationTime := time.Now().Add(10 * time.Hour)
	claims := &Claims{
		Email: 			loginCredentials.Email,
		StandardClaims: jwt.StandardClaims{
		ExpiresAt: 		expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Error creating token"})
	}

	sessionID := uuid.New().String()

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  expirationTime,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	c.SetCookie(cookie)
	UserID = foundUser.ID
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   tokenString,
		"user_id": foundUser.ID,
		"expires": expirationTime,
	})
}

func Logout(c echo.Context) error {
	_, err := getSessionAndUserID(c)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Now().Add(-24 * time.Hour),
		Path:    "/",
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Logout successful"})
}
