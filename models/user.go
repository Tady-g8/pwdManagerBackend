package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `json:"username" gorm:"unique"`
	Password  string     `json:"password"`
	Passwords []Password `gorm:"foreignKey:UserID"`
}
