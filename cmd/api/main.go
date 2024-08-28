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


	// e.GET("/users", )
	e.GET("/", handlers.HelloWorldHandler())
	// e.GET("/users", handlers.GetUsers)
	e.POST("/register", handlers.CreateUser)
	e.POST("/login", handlers.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
