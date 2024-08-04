package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Type        string      `gorm:"type:varchar(255);not null"`
	Name        string      `gorm:"type:varchar(255);not null"`
	Rarity      string      `gorm:"type:varchar(255);not null"`
	Cost        int         `gorm:"not null"`
	IsBuyable   bool        `gorm:"not null"`
	IsSellable  bool        `gorm:"not null"`
	Inventories []Inventory `gorm:"foreignKey:ItemID"`
	Auctions    []Auction   `gorm:"foreignKey:ItemID"`
}
