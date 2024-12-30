package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey"`
	Name          string    `gorm:"size:255"`
	Password     string    `gorm:"size:255"`
	Phone     	string    `gorm:"size:255;unique"`
	Birthday  time.Time
	Gender  string    `gorm:"size:255"`
	About       string    `gorm:"type:text"`
	Job     string    `gorm:"size:255"`
	Email         string    `gorm:"size:255"`
	IDProvince    string    `gorm:"size:255"`
	IDDistrict        string    `gorm:"size:255"`
	IsAdmin       bool
	Address        []Address  `gorm:"foreignKey:UserID"`
	Shop          Shop      `gorm:"foreignKey:UserID"`
	Trx     []Trx     `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}