package passwordpipeline

import (
	"github.com/Tady-g8/pwdManagerBackend/passwordpipeline/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GeneratePassword(c *fiber.Ctx, db *gorm.DB) error {

	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	appName := c.Params("appName")
	if appName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Application name is required",
		})
	}

	securePassword, err := utils.GenerateSecurePassword()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate password",
		})
	}

	salt, err := utils.GenerateSalt()
	if err != nil {
		return err
	}

	usersPassword, err := utils.GetUserMasterPassword(uint(userId), db)
	if err != nil {
		return err
	}

	encryptionKey, err := utils.GenerateEncryptionKey(usersPassword, salt)
	if err != nil {
		return err
	}

	encryptedPassword, err := utils.EncryptPassword(securePassword, encryptionKey)
	if err != nil {
		return err
	}

	err = utils.StoreEncryptedPassword(appName, encryptedPassword, userId, salt, db)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message":       "Password created successfully",
		"appName":       appName,
		"generated pwd": securePassword,
	})
}
