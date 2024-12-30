package main

import (
	"os"
	"log"
	"mini-marketplace/db"
	"mini-marketplace/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	db.RunMigration()

	// Inisialisasi JWT secret
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Setup Routes
	routes.SetupRoutes(app, db.DB, jwtSecret)
	routes.ProfileRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
