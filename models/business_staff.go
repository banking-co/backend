package models

import (
	"gorm.io/gorm"
)

type BusinessStaff struct {
	gorm.Model
	BusinessID uint     `gorm:"not null"`
	EmployerID uint     `gorm:"not null"`
	WorkerID   uint     `gorm:"not null"` // bot - 0, user - user_id
	UserType   uint8    `gorm:"not null"` // bot - 0, user - 1
	RoleID     uint8    `gorm:"not null"` // the position that the person holds
	Salary     int      `gorm:"not null"` // per hours
	Business   Business `gorm:"foreignKey:BusinessID"`
}
