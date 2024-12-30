package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"mini-marketplace/db"
	"mini-marketplace/models"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type UpdateProfileRequest struct {
	Name     string `json:"name" validate:"omitempty,max=100"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Gender   string `json:"gender" validate:"omitempty,oneof=male female"`
	About    string `json:"about" validate:"omitempty,max=500"`
	Job      string `json:"job" validate:"omitempty,max=100"`
}

func ViewProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint) // Ambil userId dari middleware otentikasi

	// Cari data pengguna berdasarkan userId
	var user models.User
	if err := db.DB.First(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Kembalikan informasi pengguna
	return c.JSON(fiber.Map{
		"message": "User profile retrieved successfully",
		"user": fiber.Map{
			"id":     user.ID,
			"email":  user.Email,
			"name":   user.Name,
			"gender": user.Gender,
			"about":  user.About,
			"job":    user.Job,
		},
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint) // Ambil userId dari middleware otentikasi
	var req UpdateProfileRequest

	// Parsing input
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validasi input menggunakan validator
	if err := validate.Struct(&req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errorMessages[fieldErr.Field()] = fieldErr.Error()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessages})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Cari user berdasarkan ID
	var user models.User
	if err := db.DB.First(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update data user
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = string(hashedPassword) // Simpan langsung untuk contoh (gunakan hashing jika diperlukan)
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	if req.About != "" {
		user.About = req.About
	}
	if req.Job != "" {
		user.Job = req.Job
	}

	// Simpan perubahan
	if err := db.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully", "user": user})
}

func DeleteProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint) // Ambil userId dari middleware otentikasi

	// Cari user berdasarkan ID
	var user models.User
	if err := db.DB.First(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Hapus user dan data terkait
	if err := db.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile deleted successfully"})
}
