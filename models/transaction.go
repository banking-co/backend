package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	UserID    uint       `gorm:"not null"`
	Type      string     `gorm:"type:varchar(50);not null"`  // types
	Status    string     `gorm:"type:varchar(100);not null"` // pending - 0, minus - 1 plus - 2
	Profit    int        `gorm:"not null;default:5"`
	ExpiredAt *time.Time `gorm:"default:null"`
	User      User       `gorm:"foreignKey:UserID"`
}
