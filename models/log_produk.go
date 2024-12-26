package models

import (
	"time"
	"gorm.io/gorm"
)

type LogProduk	struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProduk      uint      `gorm:"not null" json:"id_produk"`
	NamaProduk    string    `gorm:"type:varchar(255);not null" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255);unique;not null" json:"slug"`
	HargaReseller string    `gorm:"type:varchar(255);not null" json:"harga_reseller"`
	HargaKonsumen string    `gorm:"type:varchar(255);not null" json:"harga_konsumen"`
	Deskripsi     string    `gorm:"type:text;not null" json:"deskripsi"`
	CreatedAt     time.Time `gorm:"type:date" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:date" json:"updated_at"`
	IDToko        uint      `gorm:"not null" json:"id_toko"`     // Foreign Key ke tabel Toko
	IDKategori    uint      `gorm:"not null" json:"id_category"` // Foreign Key ke tabel Kategori

	// Relasi Many-to-One dengan Toko
	Toko Toko `gorm:"foreignKey:IDToko;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko"`

	// Relasi One-to-One dengan Kategori
	Kategori Kategori `gorm:"foreignKey:IDKategori;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"kategori"`

	// Relasi tambahan (sesuai struktur tabel sebelumnya)
	Produk     Produk       `gorm:"foreignKey:IDProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"produk"`         // One-to-One dengan Produk
	FotoProduk []FotoProduk `gorm:"foreignKey:IDLogProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"foto_produk"` // One-to-Many dengan FotoProduk
}