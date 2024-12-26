package models

import (
	"time"
	"gorm.io/gorm"
)

type DetailTransaksi struct {
	gorm.Model
	ID	uint	`gorm:"primaryKey;autoIncrement" json:"id"`
	IDTrx	Transaksi	`gorm:"foreignKey:ID" json:"id_trx"`
	IDLogProduk	LogProduk	`gorm:""`
	IDToko	Toko	`gorm:"foreignKey:ID;not null" json:"id_toko"`
	Kuantitas	int	`gorm:"not null" json:"kuantitas"`
	HargaTotal	int	`gorm:"not null" json:"harga_total"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}