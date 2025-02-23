package utils

import (
	"github.com/Tady-g8/pwdManagerBackend/models"
	"gorm.io/gorm"
)

func GetRecord(userId int, appName string, db *gorm.DB) (string, string, error) {
	var password string
	var salt string
	result := db.Model(&models.Password{}).
		Select("encrypted_password, salt").
		Where("user_id = ? AND app_name = ?", userId, appName).
		First(&models.Password{}).Scan(&struct {
		EncryptedPassword *string
		Salt              *string
	}{
		EncryptedPassword: &password,
		Salt:              &salt,
	})
	if result.Error != nil {
		return "", "", result.Error
	}
	return password, salt, nil
}
