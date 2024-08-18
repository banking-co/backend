package models

import (
	"gorm.io/gorm"
	"rabotyaga-go-backend/types"
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

func GetItems(db *gorm.DB) ([]Item, error) {
	var res []Item

	if err := db.Where("").Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func GetItemsByType(db *gorm.DB, t types.ItemType) ([]*Item, error) {
	var res []*Item

	if err := db.Where("type = ?", t).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func GetItemsByTypeAndRarity(db *gorm.DB, t types.ItemType, r types.ItemRarity) ([]*Item, error) {
	var res []*Item

	if err := db.Where("type = ? AND rarity = ?", t, r).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
