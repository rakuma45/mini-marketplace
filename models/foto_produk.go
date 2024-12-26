package models

import (
	"time"
	"gorm.io/gorm"
)

type FotoProduk struct {
	gorm.Model
	ID		uint	`gorm:"primaryKey;autoIncrement" json:"id"`
	URL		string	`json:"url"`
	IDProduk	Produk	`gorm:"foreignKey:ID;not null" json:"id_produk"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}