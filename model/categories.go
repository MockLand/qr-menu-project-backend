package model

type Categories struct {
	ID          int    `json:"category_id" gorm:"column:category_id"`
	UserId      int    `json:"user_id" gorm:"column:user_id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}
