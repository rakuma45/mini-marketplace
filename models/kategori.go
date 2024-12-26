package models

import (
	"time"
	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model
	ID	uint	`gorm:"primaryKey;autoIncrement" json:"id"`
	NamaKategori	string	`gorm:"not null" json:"nama_kategori"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Toko Toko `gorm:"foreignKey:IDKategori;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko"`
}