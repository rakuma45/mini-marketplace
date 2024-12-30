package db

import (
	"fmt"
	"log"
	"mini-marketplace/models"
)

func RunMigration() {
    err := DB.AutoMigrate(
        &models.User{},
        &models.Shop{},
        &models.ProductPhotos{},
        &models.Category{},
        &models.Product{},
        &models.Address{},
        &models.ProductLog{},
        &models.Trx{},
        &models.DetailTrx{},
    )
	if err != nil {
		log.Fatal("Gagal melakukan migrasi:", err)
	}

	fmt.Println("Migrasi tabel berhasil.")
}