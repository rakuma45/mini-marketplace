package controllers

import (
	"mini-marketplace/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB        *gorm.DB
	JWTSecret []byte
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Name         string `json:"name" validate:"required"`
		Password    string `json:"password" validate:"required"`
		Email        string `json:"email" validate:"required,email"`
		Phone       string `json:"phone" validate:"required,numeric"`
		Birthday string `json:"birthday" validate:"required"`
	}

	validate := validator.New()

	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validasi input
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Parsing TanggalLahir ke time.Time
	brithDay, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format for TanggalLahir"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Simpan ke database
	user := &models.User{
		Name:         req.Name,
		Password:    string(hashedPassword),
		Email:        req.Email,
		Phone:       req.Phone,
		Birthday: brithDay, // Menggunakan tanggal yang telah diparsing
	}
	
	if err := ac.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	store := models.Shop{ShopName: "Toko " + req.Name, UserID: user.ID}
	if err := ac.DB.Create(&store).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to automaticly add Toko"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User and Toko registered successfully"})
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Phone     string `json:"phone"`
		Password string `json:"password"`
	}

	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	var user models.User
	if err := ac.DB.Where("phone = ?", req.Phone).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Periksa kecocokan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(ac.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"message":"Login Success","token": tokenString})
}