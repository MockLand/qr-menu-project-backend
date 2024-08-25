package model

import (
)


type Users struct {
	ID        int    `json:"user_id" gorm:"column:user_id"`
	Username  string `json:"username" gorm:"column:username"`
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password_hash" gorm:"column:password_hash"`
}