package models

import (
	"gorm.io/gorm"
	"time"
)

type Ban struct {
	gorm.Model
	UserID      uint       `gorm:"not null"`
	InitiatorID uint       `gorm:"not null"`
	Until       *time.Time `gorm:"default:null"`
	Reason      string     `gorm:"type:text;default:null"`
	User        User       `gorm:"foreignKey:UserID"`
	Initiator   User       `gorm:"foreignKey:InitiatorID"`
}
