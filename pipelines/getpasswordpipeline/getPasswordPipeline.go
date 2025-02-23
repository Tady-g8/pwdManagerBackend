package getpasswordpipeline

import (
	"github.com/Tady-g8/pwdManagerBackend/pipelines/getpasswordpipeline/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPassword(c *fiber.Ctx, db *gorm.DB) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	appName := c.Params("appName")
	if appName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid app name format",
		})
	}

	encryptedPassword, salt, err := utils.GetRecord(userId, appName, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid app name format",
		})
	}

	userMaster, err := utils.GetUserMasterPassword(userId, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid app name format",
		})
	}

	decryptionKey, err := utils.GenerateEncryptionKey(userMaster, salt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid app name format",
		})
	}

	decryptedPassword, err := utils.DecryptPassword(encryptedPassword, decryptionKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid app name format",
		})
	}

	return c.JSON(fiber.Map{
		"message":       "Password decrypted successfully",
		"decrypted pwd": decryptedPassword,
	})
}
