package models

import (
	"time"
	"gorm.io/gorm"
)

type ProductPhotos struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	ProductID  uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductID"`
	URL       string    `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}