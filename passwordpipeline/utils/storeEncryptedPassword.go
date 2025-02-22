package utils

import (
	"fmt"

	"github.com/Tady-g8/pwdManagerBackend/models"
	"gorm.io/gorm"
)

func StoreEncryptedPassword(appName string, encryptedPassword string, userId int, salt string, db *gorm.DB) error {

	password := models.Password{
		AppName: appName,
		Value:   encryptedPassword,
		Salt:    salt,
		UserID:  uint(userId),
	}

	result := db.Create(&password)
	if result.Error != nil {
		return fmt.Errorf("failed to store password: %v", result.Error)
	}

	return nil
}
