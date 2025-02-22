package passwordpipeline

import (
	"fmt"
	passwordPipelineUtils "server/generatePassword/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PasswordPipeline struct {
	db *gorm.DB
}

func NewPasswordPipeline(db *gorm.DB) *PasswordPipeline {
	return &PasswordPipeline{
		db: db,
	}
}

func (p *PasswordPipeline) GeneratePassword(c *fiber.Ctx) error {

	fmt.Print("called GeneratePassword")

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

	securePassword, err := passwordPipelineUtils.GenerateSecurePassword()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate password",
		})
	}

	salt, err := passwordPipelineUtils.GenerateSalt()
	if err != nil {
		return err
	}

	usersPassword, err := passwordPipelineUtils.GetUserMasterPassword(p.db, uint(userId))
	if err != nil {
		return err
	}

	encryptionKey, err := passwordPipelineUtils.GenerateEncryptionKey(usersPassword, salt)
	if err != nil {
		return err
	}

	encryptedPassword, err := passwordPipelineUtils.EncryptPassword(securePassword, encryptionKey)
	if err != nil {
		return err
	}

	err = passwordPipelineUtils.StoreEncryptedPassword(p.db, appName, encryptedPassword, userId, salt)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Password created successfully",
		"appName": appName,
	})
}
