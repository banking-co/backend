package models

import (
	"gorm.io/gorm"
	"time"
)

type Inventory struct {
	gorm.Model
	ItemID   uint       `gorm:"not null"`
	BuyerID  uint       `gorm:"not null"`
	SellerID uint       `gorm:"not null"`
	BuyAt    time.Time  `gorm:"autoCreateTime"`
	SellAt   *time.Time `gorm:"default:null"`
	Item     Item       `gorm:"foreignKey:ItemID"`
	Buyer    User       `gorm:"foreignKey:BuyerID"`
	Seller   User       `gorm:"foreignKey:SellerID"`
}
