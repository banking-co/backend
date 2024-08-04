package models

import (
	"gorm.io/gorm"
)

type BusinessProfit struct {
	gorm.Model
	BusinessID uint     `gorm:"not null"`
	UserID     uint     `gorm:"not null"`
	Business   Business `gorm:"foreignKey:BusinessID"`
	User       User     `gorm:"foreignKey:UserID"`
}
