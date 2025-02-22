package utils

import (
	"fmt"

	"github.com/Tady-g8/pwdManagerBackend/models"

	"gorm.io/gorm"
)

func GetUserMasterPassword(userId uint, db *gorm.DB) (string, error) {

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
