package models

import (
	"gorm.io/gorm"
)

type BusinessRole struct {
	gorm.Model
	UserID     uint     `gorm:"not null"`
	BusinessID uint     `gorm:"not null"`
	RoleId     uint8    `gorm:"not null"`
	RoleName   string   `gorm:"type:varchar(255);not null"`
	User       User     `gorm:"foreignKey:UserID"`
	Business   Business `gorm:"foreignKey:BusinessID"`
}
