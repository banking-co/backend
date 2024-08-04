package models

import (
	"gorm.io/gorm"
	"time"
)

type BusinessUpgrade struct {
	gorm.Model
	Name       string     `gorm:"type:tinytext;not null"`
	Type       int        `gorm:"not null"`
	BusinessID uint       `gorm:"not null"`
	BuyAt      *time.Time `gorm:"default:null"`
	SellAt     *time.Time `gorm:"default:null"`
	Business   Business   `gorm:"foreignKey:BusinessID"`
}
