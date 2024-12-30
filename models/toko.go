package models

import (
	"time"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	ShopName   string    `gorm:"size:255"`
	UrlPhoto    string    `gorm:"size:255"`
	Product     []Product  `gorm:"foreignKey:ShopID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}