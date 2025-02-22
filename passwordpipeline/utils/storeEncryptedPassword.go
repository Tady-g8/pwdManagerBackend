package utils

import (
	"fmt"
	"log"

	"github.com/Tady-g8/pwdManagerBackend/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func StoreEncryptedPassword(appName string, encryptedPassword string, userId int, salt string) error {
	db, err := gorm.Open(sqlite.Open("../../passwords.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

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
