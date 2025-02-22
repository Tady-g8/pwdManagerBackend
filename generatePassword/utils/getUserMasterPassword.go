package utils

import (
	"fmt"
	"server/models"

	"gorm.io/gorm"
)

func GetUserMasterPassword(db *gorm.DB, userId uint) (string, error) {
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
