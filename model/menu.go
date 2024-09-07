package model

type Menus struct {
	ID         int    `json:"menu_id" gorm:"column:menu_id"`
	User_id    int    `json:"user_id" gorm:"column:user_id"`
	Name       string `json:"name" gorm:"column:name"`
	Start_date string `json:"start_date" gorm:"column:start_date"`
	End_date   string `json:"end_date" gorm:"column:end_date"`
}

type UpdateMenuCredentials struct {
	User_id string `json:"user_id" gorm:"column:user_id"`
	Name       string `json:"name" binding:"required"`
	Start_date string `json:"start_date" binding:"required"`
	End_date   string `json:"end_date" binding:"required"`
}