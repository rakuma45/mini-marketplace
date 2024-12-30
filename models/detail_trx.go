package models

import (
	"time"
	"gorm.io/gorm"
)

type DetailTrx struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey"`
	TrxID        uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TrxID"`
	ProductLogID  uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:LogProductID"`
	ShopID       uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ShopID"`
	Qty    		int
	TotalPrice   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}