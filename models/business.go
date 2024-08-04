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
	Profits  []BusinessProfit  `gorm:"foreignKey:BusinessID"`
}
