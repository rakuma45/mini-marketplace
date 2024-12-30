package models

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID           uint         `gorm:"primaryKey"`
	ProductName   string       `gorm:"size:255"`
	Slug         string       `gorm:"size:255"`
	ResellerPrice string      `gorm:"size:255"`
	ConsumerPrice string      `gorm:"size:255"`
	Stock         int
	Description    string       `gorm:"type:text"`
	ShopID       uint         `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ShopID"`
	CategoryID   uint         `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID"`
	ProductPhotos   []ProductPhotos `gorm:"foreignKey:ProductID"`
	ProductLog    ProductLog    `gorm:"foreignKey:ProductID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}