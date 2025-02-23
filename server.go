package main

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Tady-g8/pwdManagerBackend/models"
	"github.com/Tady-g8/pwdManagerBackend/utils"
)

func main() {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Password{})
	if err != nil {
		panic("failed to migrate database")
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/createPwd/:userId/:appName", func(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		return createPasswordpipeline.GeneratePassword(c, db)
	})

	app.Get("/getAppNames/:userId", func(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		return utils.GetUsersAppNames(c, db)
	})

	app.Get("/getPassword/:userId/:appName", func(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		return getPasswordPipeline.GetPassword(c, db)
	})

	log.Fatal(app.Listen(":3000"))
}
