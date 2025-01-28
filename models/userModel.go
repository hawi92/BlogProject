package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Name string `json:"name"`
	Bio    string `json:"bio"`
	Role string `json:"role"`
}

