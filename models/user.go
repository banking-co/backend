package models

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
)

type User struct {
	gorm.Model
	Username          string          `gorm:"type:varchar(255);not null;unique"`
	Bans              []Ban           `gorm:"foreignKey:UserID"`
	Balances          []Balance       `gorm:"foreignKey:UserID"`
	Bonuses           []Bonus         `gorm:"foreignKey:UserID"`
	Businesses        []Business      `gorm:"foreignKey:UserID"`
	BusinessStaff     []BusinessStaff `gorm:"foreignKey:WorkerID"`
	ReferralsSent     []Referral      `gorm:"foreignKey:ReferrerID"`
	ReferralsReceived []Referral      `gorm:"foreignKey:ReferralID"`
	Transactions      []Transaction   `gorm:"foreignKey:UserID"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	balances := []Balance{
		{UserID: u.ID, Currency: "donate", Amount: 0},
		{UserID: u.ID, Currency: "usd", Amount: 1000},
		{UserID: u.ID, Currency: "btc", Amount: 1},
	}

	if err := tx.Create(&balances).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(db *gorm.DB, username int) (*User, error) {
	var user User
	uid := "id" + strconv.Itoa(username)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Preload("Bans").
			Preload("Balances").
			Where("username = ?", uid).
			First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				user = User{
					Username: uid,
				}

				if err := tx.Create(&user).Error; err != nil {
					return err
				}
			}

			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}
