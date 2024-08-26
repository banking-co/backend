package models

import (
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	UserID   uint              `gorm:"not null"`
	Name     string            `gorm:"type:varchar(255);not null"`
	User     User              `gorm:"foreignKey:UserID"`
	Upgrades []BusinessUpgrade `gorm:"foreignKey:BusinessID"`
	Staff    []BusinessStaff   `gorm:"foreignKey:BusinessID"`
	Roles    []BusinessRole    `gorm:"foreignKey:BusinessID"`
	Profits  []BusinessProfit  `gorm:"foreignKey:BusinessID"`
}

func GetBusinessById(db *gorm.DB, bid uint) (*Business, error) {
	var business Business

	if err := db.
		Preload("User").
		Preload("Upgrades").
		Preload("Staff").
		Preload("Profits").
		Where(`id = ?`, bid).
		First(&business).Error; err != nil {
		return nil, err
	}

	return &business, nil
}

func GetBusinessByUserId(db *gorm.DB, uid uint) (*Business, error) {
	var business Business

	if err := db.
		Preload("User").
		Preload("Upgrades").
		Preload("Staff").
		Preload("Profits").
		Preload("Roles").
		Where(`user_id = ?`, uid).
		First(&business).Error; err != nil {
		return nil, err
	}

	return &business, nil
}
