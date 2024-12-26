package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID				int			`gorm:"primaryKey" json:"id"`
	Nama			string 		`gorm:"not_nill" json:"nama"`
	KataSandi   	string    	`gorm:"not null" json:"kata_sandi"`
	NoTelp      	string    	`gorm:"unique;not null" json:"no_telp"`
	TanggalLahir 	time.Time 	`gorm:"type:date;not null" json:"tanggal_lahir"`
	JenisKelamin 	string    	`gorm:"not null" json:"jenis_kelamin"`
	Tentang     	string    	`json:"tentang"`
	Pekerjaan   	string    	`json:"pekerjaan"`
	Email       	string    	`gorm:"not null" json:"email"`
	IDProvinsi  	string    	`json:"id_provinsi"`
	IDKota      	string    	`json:"id_kota"`
	IsAdmin     	bool      	`gorm:"default:false" json:"is_admin"`
	CreatedAt   	time.Time 	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   	time.Time 	`gorm:"autoUpdateTime" json:"updated_at"`

	Toko     Toko   `gorm:"foreignKey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"toko"` // Relasi One-to-One dengan Toko
	Alamat 			[]Alamat	`gorm:"foreignKey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"alamat"`
}