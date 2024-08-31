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



type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}


type JWTOutput struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expires"`
}

type SessionData struct {
	Token  string    `json:"token"`
	UserId uuid.UUID `json:"user_id"`
}



func Login(c echo.Context) error {
	var loginCredentials model.LoginCredentials

	if err := c.Bind(&loginCredentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	foundUser := model.Users{}
	result := database.DB.Where("email = ?", loginCredentials.Email).First(&foundUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Database error"})
	}

	if foundUser.Password != loginCredentials.Password {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid email or password"})
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: loginCredentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   tokenString,
		"user_id": foundUser.ID,
		"expires": expirationTime,
	})
}


