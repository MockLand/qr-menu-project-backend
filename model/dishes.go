package model


type Dishes struct {
	ID           int    `json:"dish_id" gorm:"column:dish_id"`
	UserId int `json:"user_id" gorm:"column:user_id"`
	Name string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Price float64 `json:"price" gorm:"column:price"`
	CategoryId int `json:"category_id" gorm:"column:category_id"`
	IsAvailable bool `json:"is_"`
	
}

