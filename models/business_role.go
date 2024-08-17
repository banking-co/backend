package models

import (
	"gorm.io/gorm"
)

type BusinessRole struct {
	gorm.Model
	BusinessID uint     `gorm:"not null"`
	RoleId     uint8    `gorm:"not null"`
	RoleName   string   `gorm:"type:varchar(255);not null"`
	Business   Business `gorm:"foreignKey:BusinessID"`
}
