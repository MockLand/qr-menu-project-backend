package model

type Ingredient struct {
	ID           int    `json:"ingredient_id" gorm:"column:ingredient_id"`
	UserId       int    `json:"user_id" gorm:"column:user_id"`
	Name         string `json:"name" gorm:"column:name"`
	AllergenInfo string `json:"allergen_info" gorm:"column:allergen_info"`
}

type DishIngredients struct {
	DishId       int `json:"dish_id" gorm:"column:dish_id"`
	IngredientId int `json:"ingredient_id" gorm:"column:ingredient_id"`
	UserId      int `json:"user_id" gorm:"column:user_id"`
	Quantity     string `json:"quantity" gorm:"column:quantity"`
}
