package models

import (
	"time"
	"gorm.io/gorm"
)

type Trx struct {
	gorm.Model
	ID               uint         `gorm:"primaryKey"`
	UserID           uint         `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	ShippingAddress uint         `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ID"`
	TotalPrice       int
	InvoiceID      string       `gorm:"size:255"`
	Payment      string       `gorm:"size:255"`
	DetailTrx  []DetailTrx  `gorm:"foreignKey:TrxID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}	