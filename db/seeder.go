package db

import (	
	"log"
	"time"
	"mini-marketplace/models"
)

// Helper function to parse date string to time.Time
func parseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Fatal("Error parsing date:", err)
	}
	return date
}

func SeedData() {
	// Seeder untuk data User
	users := []models.User{
		{
			Name:          "John Doe",
			Password:     "password123",
			Phone:        "08123456789",
			Birthday:  parseDate("1995-05-10"),
			Gender:  "male",
			Email:         "johndoe@example.com",
		},
		{
			Name:          "Jane Smith",
			Password:     "password123",
			Phone:        "08198765432",
			Birthday:  parseDate("1995-05-10"),
			Gender:  "female",
			Email:         "janesmith@example.com",
		},
	}

	for _, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			log.Println("Gagal menambahkan user:", err)
		}
	}

	// Seeder untuk data Toko
	toko := []models.Shop{
		{
			ShopName:      "Toko Sukses Jaya",
			UserID: 1, // ID User sebagai pemilik toko
		},
		{
			ShopName:      "Toko Surya Abadi",
			UserID: 2, // ID User sebagai pemilik toko
		},
	}

	if err := DB.Create(&toko).Error; err != nil {
		log.Println("Gagal menambahkan toko:", err)
	}

	// Seeder untuk data Produk
	produk := []models.Product{
		{
			ProductName:        "Sabun Cuci",
			Description:   "Sabun cuci berkualitas tinggi",
			ConsumerPrice:       "10000",
			ResellerPrice:		 "9500",
			Stock:        50,
			ShopID:      1, // ID Toko
		},
		{
			ProductName:        "Gula Pasir",
			Description:   "Gula pasir murni",
			ConsumerPrice:       "7000",
			ResellerPrice:		 "6500",
			Stock:        30,
			ShopID:      2, // ID Toko
		},
	}

	for _, p := range produk {
		if err := DB.Create(&p).Error; err != nil {
			log.Println("Gagal menambahkan produk:", err)
		}
	}

	log.Println("Seeder data berhasil dijalankan!")
}