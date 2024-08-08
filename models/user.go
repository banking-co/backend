package models

import (
	"errors"
	"github.com/SevereCloud/vksdk/v3/api"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/vk"
	"strconv"
)

type User struct {
	gorm.Model
	VkId              int             `gorm:"not null"`
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
		{Importance: 0, UserID: u.ID, Currency: "usd", Amount: 50000},
		{Importance: 1, UserID: u.ID, Currency: "btc", Amount: 100},
		{Importance: 2, UserID: u.ID, Currency: "donate", Amount: 0},
	}

	if err := tx.Create(&balances).Error; err != nil {
		return err
	}

	business := Business{
		UserID: u.ID,
		Name:   "Bank",
	}

	if err := tx.Create(&business).Error; err != nil {
		return err
	}

	businessRole := BusinessRole{
		UserID:     u.ID,
		BusinessID: business.ID,
		RoleId:     types.BusinessRoleBot,
		RoleName:   "bot",
	}

	if err := tx.Create(&businessRole).Error; err != nil {
		return err
	}

	businessStaff := BusinessStaff{
		BusinessID: business.ID,
		EmployerID: 0,
		UserType:   0,
		WorkerID:   u.ID,
		RoleID:     types.BusinessRoleBot,
		Salary:     0,
	}

	if err := tx.Create(&businessStaff).Error; err != nil {
		return err
	}

	return nil
}

func GetVkUserInfo(id int) (*api.UsersGetResponse, error) {
	users, err := vk.Api.UsersGet(api.Params{
		"user_ids": id,
	})

	if err != nil {
		return nil, err
	}

	return &users, nil
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
					VkId:     username,
					Username: uid,
				}

				if err := tx.Create(&user).Error; err != nil {
					return err
				}

				return nil
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
