package models

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID           uint       `gorm:"primaryKey"`
	CategoryName string     `gorm:"size:255"`
	Product       []Product   `gorm:"foreignKey:CategoryID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}