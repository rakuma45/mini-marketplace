package db

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

