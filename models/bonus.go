package models

import (
	"gorm.io/gorm"
	"time"
)

type Bonus struct {
	gorm.Model
	UserID     uint       `gorm:"not null"`
	Type       string     `gorm:"type:varchar(255);not null"`
	ReceivedAt *time.Time `gorm:"default:null"`
	User       User       `gorm:"foreignKey:UserID"`
}
