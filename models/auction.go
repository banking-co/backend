package models

import (
	"gorm.io/gorm"
	"time"
)

type Auction struct {
	gorm.Model
	ItemID   uint      `gorm:"not null"`
	SellerID uint      `gorm:"not null"`
	BuyerID  uint      `gorm:"not null"`
	SellAt   time.Time `gorm:"autoCreateTime"`
	Item     Item      `gorm:"foreignKey:ItemID"`
	Seller   User      `gorm:"foreignKey:SellerID"`
	Buyer    User      `gorm:"foreignKey:BuyerID"`
}
