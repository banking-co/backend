package models

import (
	"gorm.io/gorm"
)

type Referral struct {
	gorm.Model
	ReferrerID uint `gorm:"not null"`
	ReferralID uint `gorm:"not null"`
	Profit     int  `gorm:"not null;default:5"`
	Referrer   User `gorm:"foreignKey:ReferrerID"`
	Referral   User `gorm:"foreignKey:ReferralID"`
}
