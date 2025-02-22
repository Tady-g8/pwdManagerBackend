package main

import (
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"github.com/Tady-g8/pwdManagerBackend/passwordpipeline"
)

func main() {

	fmt.Print("hello???")

	db, err := gorm.Open(sqlite.Open("passwords.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/createPwd/:userId/:appName", passwordpipeline.GeneratePassword)

	log.Fatal(app.Listen(":3000"))
}
