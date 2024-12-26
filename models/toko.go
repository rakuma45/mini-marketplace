package models

import (
	"time"
	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser    uint      `gorm:"not null" json:"id_user"`
	NamaToko  string    `gorm:"type:varchar(255);not null" json:"nama_toko"`
	UrlFoto   string    `gorm:"type:varchar(255)" json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Produk       []Produk         `gorm:"foreignKey:IDToko" json:"produk"` // Relasi One-to-Many dengan Produk
	LogProduk    []LogProduk      `gorm:"foreignKey:IDToko" json:"log_produk"` // Relasi One-to-Many dengan LogProduk
	DetailTrx    []DetailTransaksi `gorm:"foreignKey:IDToko" json:"detail_trx"` // Relasi One-to-Many dengan DetailTransaksi
	User         *User             `gorm:"foreignKey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"` // Relasi One-to-One dengan User
}