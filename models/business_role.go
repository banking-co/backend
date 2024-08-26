package models

import (
	"errors"
	"gorm.io/gorm"
)

type BusinessRole struct {
	gorm.Model
	BusinessID uint     `gorm:"not null"`
	RoleId     uint8    `gorm:"not null"`
	RoleName   string   `gorm:"type:varchar(255);not null"`
	Business   Business `gorm:"foreignKey:BusinessID"`
}

func GetBusinessRolesByBusinessId(db *gorm.DB, bid uint) ([]BusinessRole, error) {
	var businessRoles []BusinessRole

	if bid < 1 {
		return nil, errors.New("business id is nil")
	}

	if err := db.
		Where(`business_id = ?`, bid).
		Find(&businessRoles).Error; err != nil {
		return nil, err
	}

	return businessRoles, nil
}
