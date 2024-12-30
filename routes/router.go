package routes

import (
	"github.com/gofiber/fiber/v2"
	"mini-marketplace/controllers"
	"mini-marketplace/middlewares"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwtSecret []byte) {
	authController := controllers.AuthController{
		DB:        db,
		JWTSecret: jwtSecret,
	}

	// Setup routes
	app.Get("/", func(c *fiber.Ctx) error {
  		return c.SendString("Hello, World!")
 	})

	api := app.Group("/api")
	api.Post("/register", authController.Register)
	api.Post("/login", authController.Login)
}

func ProfileRoutes(app *fiber.App) {
	api := app.Group("/api")
	profile := api.Group("/profile", middlewares.AuthMiddleware)

	profile.Get("/", controllers.ViewProfile)   // Melihat informasi akun
	profile.Put("/", controllers.UpdateProfile) // Mengubah informasi akun
	profile.Delete("/", controllers.DeleteProfile) // Menghapus akun
}