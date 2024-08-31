package model

import (
	"time"
)

type Users struct {
	ID           int       `json:"user_id" gorm:"column:user_id"`
	Username     string    `json:"username" gorm:"column:username"`
	Email        string    `json:"email" gorm:"column:email"`
	Password     string    `json:"password_hash" gorm:"column:password_hash"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type LoginCredentials struct {
    Email    string `json:"email" gorm:"column:email"`
    Password string `json:"password_hash" gorm:"column:password_hash"`
}
