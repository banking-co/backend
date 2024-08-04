package models

import (
	"gorm.io/gorm"
)

type PromoCode struct {
	gorm.Model
	Code   string `gorm:"type:varchar(30);not null"`
	Count  int    `gorm:"not null"`
	Type   string `gorm:"type:varchar(50);not null"` // promocode type: money, item, etc
	Profit int    `gorm:"not null;default:1"`
}
