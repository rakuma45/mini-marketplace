package models

import (
	"time"
	"gorm.io/gorm"
)

type Alamat struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser       User      `gorm:"foreignKey:ID;not null" json:"id_user"`
	JudulAlamat  string    `gorm:"not null" json:"judul_alamat"`
	NamaPenerima string    `gorm:"not null" json:"nama_penerima"`
	NoTelp       string    `gorm:"not null" json:"no_telp"`
	DetailAlamat string    `gorm:"not null" json:"detail_alamat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}