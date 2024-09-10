package model

type Menus struct {
	ID         int    `json:"menu_id" gorm:"column:menu_id"`
	UserId    int    `json:"user_id" gorm:"column:user_id"`
	Name       string `json:"name" gorm:"column:name"`
	StartDate string `json:"start_date" gorm:"column:start_date"`
	EndDate   string `json:"end_date" gorm:"column:end_date"`
}

type UpdateMenuCredentials struct {
	UserId string `json:"user_id" gorm:"column:user_id"`
	Name       string `json:"name" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}