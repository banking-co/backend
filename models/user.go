package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string          `gorm:"type:varchar(255);not null"`
	Bans              []Ban           `gorm:"foreignKey:UserID"`
	Bonuses           []Bonus         `gorm:"foreignKey:UserID"`
	Businesses        []Business      `gorm:"foreignKey:UserID"`
	BusinessStaff     []BusinessStaff `gorm:"foreignKey:WorkerID"`
	ReferralsSent     []Referral      `gorm:"foreignKey:ReferrerID"`
	ReferralsReceived []Referral      `gorm:"foreignKey:ReferralID"`
	Transactions      []Transaction   `gorm:"foreignKey:UserID"`
}
