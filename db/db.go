package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mini-marketplace/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Konfigurasi koneksi database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),   // Username database
		os.Getenv("DB_PASS"),   // Password database
		os.Getenv("DB_HOST"),   // Host database
		os.Getenv("DB_PORT"),   // Port database
		os.Getenv("DB_NAME"),   // Nama database
	)

	// Membuka koneksi ke database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database:", err)
	}

	DB = database
	fmt.Println("Koneksi ke database berhasil.")
}


func RunMigration(db *gorm.DB) {
    err := DB.AutoMigrate(
        &models.User{},
        &models.Toko{},
        &models.Produk{},
        &models.Kategori{},
        &models.Alamat{},
        &models.FotoProduk{},
        &models.LogProduk{},
        &models.Transaksi{},
        &models.DetailTransaksi{},
    )
	if err != nil {
		log.Fatal("Gagal melakukan migrasi:", err)
	}

	fmt.Println("Migrasi tabel berhasil.")
}

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
			Nama:          "John Doe",
			KataSandi:     "password123",
			NoTelp:        "08123456789",
			TanggalLahir:  parseDate("1995-05-10"),
			JenisKelamin:  "Laki-laki",
			Email:         "johndoe@example.com",
		},
		{
			Nama:          "Jane Smith",
			KataSandi:     "password123",
			NoTelp:        "08198765432",
			TanggalLahir:  parseDate("1995-05-10"),
			JenisKelamin:  "Perempuan",
			Email:         "janesmith@example.com",
		},
	}

	for _, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			log.Println("Gagal menambahkan user:", err)
		}
	}

	// Seeder untuk data Toko
	toko := []models.Toko{
		{
			NamaToko:      "Toko Sukses Jaya",
			IDUser: 1, // ID User sebagai pemilik toko
		},
		{
			NamaToko:      "Toko Surya Abadi",
			IDUser: 2, // ID User sebagai pemilik toko
		},
	}

	if err := DB.Create(&toko).Error; err != nil {
		log.Println("Gagal menambahkan toko:", err)
	}

	// Seeder untuk data Produk
	produk := []models.Produk{
		{
			NamaProduk:        "Sabun Cuci",
			Deskripsi:   "Sabun cuci berkualitas tinggi",
			HargaKonsumen:       "10000",
			HargaReseller:		 "9500",
			Stok:        50,
			IDToko:      1, // ID Toko
		},
		{
			NamaProduk:        "Gula Pasir",
			Deskripsi:   "Gula pasir murni",
			HargaKonsumen:       "7000",
			HargaReseller:		 "6500",
			Stok:        30,
			IDToko:      2, // ID Toko
		},
	}

	for _, p := range produk {
		if err := DB.Create(&p).Error; err != nil {
			log.Println("Gagal menambahkan produk:", err)
		}
	}

	log.Println("Seeder data berhasil dijalankan!")
}