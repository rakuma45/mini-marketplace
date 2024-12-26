package models

import (
	"time"
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduk   string    `gorm:"not null" json:"nama_produk"`
	Slug         string    `gorm:"unique;not null" json:"slug"`
	HargaReseller string   `gorm:"not null" json:"harga_reseller"`
	HargaKonsumen string   `gorm:"not null" json:"harga_konsumen"`
	Stok         int       `gorm:"not null" json:"stok"`
	Deskripsi    string    `json:"deskripsi"`
	IDToko       uint      `gorm:"not null" json:"id_toko"`
	IDKategori   Kategori      `gorm:"foreignKey:ID;not null" json:"id_kategori"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Toko Toko `gorm:"foreignKey:IDToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"toko"`
	FotoProduk []FotoProduk `gorm:"foreignKey:IDProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"foto_produk"`
}