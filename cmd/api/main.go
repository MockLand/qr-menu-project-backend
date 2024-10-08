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
	e.GET("/menus", handlers.GetMenus)
	e.GET("/menu/:id", handlers.GetMenu)
	e.GET("/categories", handlers.GetCategories)
	e.GET("/category/:id", handlers.GetCategory)
	e.GET("/dishes", handlers.GetDishes)
	e.GET("/dish/:id", handlers.GetDish)
	e.GET("/ingredients", handlers.GetIngredients)
	e.GET("/ingredients/:id", handlers.GetIngredient)
	e.GET("/dishingredients/:dish_id", handlers.GetDishIngredients)





	e.POST("/menu", handlers.CreateMenu)
	e.POST("/register", handlers.CreateUser)
	e.POST("/login", handlers.Login)
	e.POST("/logout", handlers.Logout)
	e.POST("/category", handlers.CreateCategory)
	e.POST("/dish", handlers.CreateDish)
	e.POST("/ingredient", handlers.CreateIngredient)
	e.POST("/dishingredient", handlers.CreateDishIngredients)

	e.DELETE("/menu/:id", handlers.DeleteMenu)
	e.DELETE("/category/:id", handlers.DeleteCategory)
	e.DELETE("/dish/:id", handlers.DeleteDish)
	e.DELETE("/ingredient/:id", handlers.DeleteIngredient)
	e.DELETE("/dishingredient/:dish_id/:ingredient_id", handlers.DeleteDishIngredients)

	e.PUT("/menu/:id", handlers.UpdateMenu)
	e.PUT("/category/:id", handlers.UpdateCategory)
	e.PUT("/dish/:id", handlers.UpdateDish)
	e.PUT("/ingredient/:id", handlers.UpdateIngredient)

	e.Logger.Fatal(e.Start(":8080"))
}
