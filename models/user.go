package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/object"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"rabotyaga-go-backend/redisdb"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/vk"
	"strconv"
	"time"
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
		Name:   "Amazing Bank",
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

func GetVkUserInfo(id int) (*object.UsersUser, error) {
	ids := []int{id}
	users, err := GetVkUsersInfo(ids)

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New("GetVkUserInfo is nil")
	}

	if len(*users) > 1 {
		return nil, errors.New("GetVkUserInfo is big")
	}

	firstUser := (*users)[0]
	return &firstUser, nil
}

func GetVkUsersInfo(ids []int) (*[]object.UsersUser, error) {
	var strIds = make([]string, 0, len(ids))
	for _, id := range ids {
		strIds = append(strIds, fmt.Sprintf("vk_id%v", strconv.Itoa(id)))
	}

	usersInfoWithBytes, err := redisdb.GetAll(strIds)
	if !errors.Is(err, nil) {
		fmt.Println("GetVkUsersInfo redis error:", err)
		return nil, err
	}

	usersInfo := make([]object.UsersUser, 0, len(ids))
	if usersInfoWithBytes != nil {
		for _, u := range *usersInfoWithBytes {
			if u == nil {
				return nil, errors.New("GetVkUsersInfo user is nil")
			}

			user := object.UsersUser{}

			err := json.Unmarshal(u, &user)
			if err != nil {
				return nil, err
			}

			usersInfo = append(usersInfo, user)
		}

		if len(usersInfo) == len(ids) {
			return &usersInfo, nil
		}
	}

	existingIDs := make(map[int]bool, len(usersInfo))
	for _, user := range usersInfo {
		existingIDs[user.ID] = true
	}

	missingIDs := make([]int, 0, len(ids)-len(usersInfo))
	for _, id := range ids {
		if !existingIDs[id] {
			missingIDs = append(missingIDs, id)
		}
	}

	vkUsers, err := vk.Api.UsersGet(api.Params{
		"user_ids": missingIDs,
	}.Lang(object.LangRU))
	if err != nil {
		return nil, err
	}

	saveData := make(map[string][]byte)
	users := make([]object.UsersUser, 0, len(missingIDs))
	for _, u := range vkUsers {
		users = append(users, u)

		uB, err := json.Marshal(u)
		if err != nil {
			return nil, err
		}
		saveData[fmt.Sprintf("vk_id%v", strconv.Itoa(u.ID))] = uB
	}

	err = redisdb.SetAll(saveData, 23*time.Hour)
	if err != nil {
		return nil, err
	}

	cS := append(usersInfo, users...)

	return &cS, nil
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

func GetUserById(db *gorm.DB, id uint) (*User, error) {
	var user User

	if err := db.
		Preload("Bans").
		Preload("Balances").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, errors.New("user is nil")
	}

	return &user, nil
}
