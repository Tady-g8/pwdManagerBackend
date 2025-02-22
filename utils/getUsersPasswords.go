package utils

import (
	"github.com/Tady-g8/pwdManagerBackend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsersAppNames(c *fiber.Ctx, db *gorm.DB) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	var appNames []string

	result := db.Model(&models.Password{}).
		Where("user_id = ?", userId).
		Pluck("app_name", &appNames)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve app names",
		})
	}

	return c.JSON(appNames)
}
