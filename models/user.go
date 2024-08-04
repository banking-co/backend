package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"
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

func GetUserByUsername(db *gorm.DB, username int) (*User, error) {
	var user User
	var uid = "id" + strconv.Itoa(username)

	if err := db.
		Preload("Bans").
		Where("username = ?", uid).
		First(&user).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			user := User{
				Username: uid,
			}
			if err := db.Create(&user).Error; err != nil {
				log.Fatal("Failed to create user:", err)
				return nil, err
			} else {
				return &user, nil
			}
		} else {
			log.Fatal("Failed to query user:", err)
			return nil, err
		}
	}

	return &user, nil
}
