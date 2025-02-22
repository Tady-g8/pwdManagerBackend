package utils

import (
	"fmt"
	"log"

	"github.com/Tady-g8/pwdManagerBackend/models"
	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

func GetUserMasterPassword(userId uint) (string, error) {

	db, err := gorm.Open(sqlite.Open("../../passwords.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	var user models.User
	result := db.First(&user, userId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("failed to fetch user: %v", result.Error)
	}

	return user.Password, nil
}
