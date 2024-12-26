package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
	"mini-marketplace/db"
	"github.com/joho/godotenv"

)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()

	// Setup routes
	app.Get("/", func(c fiber.Ctx) error {
  		return c.SendString("Hello, World!")
 	})

	// Start server
	log.Fatal(app.Listen(":8080"))
}
