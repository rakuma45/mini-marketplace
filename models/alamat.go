package models

import (
	"time"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey"`
	UserID         uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	Address    string    `gorm:"size:255"`
	Receiver   string    `gorm:"size:255"`
	Phone         string    `gorm:"size:255"`
	AddressDetail   string    `gorm:"size:255"`
	Trx      []Trx     `gorm:"foreignKey:ShippingAddress"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}