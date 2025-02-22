package main

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Tady-g8/pwdManagerBackend/models"
	"github.com/Tady-g8/pwdManagerBackend/passwordpipeline"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/createPwd/:userId/:appName", passwordpipeline.GeneratePassword)

	log.Fatal(app.Listen(":3000"))
}
