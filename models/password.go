package models

import (
	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	AppName string `json:"appName"`
	Value   string `json:"value"`
	Salt    string `json:"salt"`
	UserID  uint
}
