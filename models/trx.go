package models

import (
	"time"
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser		User	`gorm:"foreignKey:ID" json:"id_user"`
	AlamatPengiriman	Alamat	`gorm:"foreignKey:ID" json:"alamat_pengiriman"`
	HargaTotal int		`json:"harga_total"`
	KodeInvoice	string	`json:"kode_invoice"`
	MethodBayar	string	`json:"method_bayar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Toko Toko `gorm:"foreignKey:IDToko;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko"`
}	