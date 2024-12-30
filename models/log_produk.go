package models

import (
	"time"
	"gorm.io/gorm"
)

type ProductLog	struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey"`
	ProductID     uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductID"`
	ProductName   string    `gorm:"size:255"`
	Slug         string    `gorm:"size:255"`
	ResellerPrice string   `gorm:"size:255"`
	ConsumerPrice string   `gorm:"size:255"`
	Description    string    `gorm:"type:text"`
	ShopID       uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ShopID"`
	CategoryID   uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}