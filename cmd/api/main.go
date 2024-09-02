package main

import (
	"qr-menu-project-backend/database"
	"qr-menu-project-backend/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e:= echo.New()

	database.NewDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.GET("/", handlers.HelloWorldHandler())


	e.POST("/menu", handlers.CreateMenu)
	e.POST("/register", handlers.CreateUser)
	e.POST("/login", handlers.Login)
	e.POST("/logout", handlers.Logout)

	e.Logger.Fatal(e.Start(":8080"))
}
