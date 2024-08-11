package models

import (
	"errors"
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

func GetBusinessStaffByBusinessId(db *gorm.DB, bid *int) ([]*BusinessStaff, error) {
	var businessStaff []*BusinessStaff

	if bid == nil {
		return nil, errors.New("business id is nil")
	}

	if err := db.
		Where(`business_id = ?`, bid).
		Find(&businessStaff).Error; err != nil {
		return nil, err
	}

	return businessStaff, nil
}
