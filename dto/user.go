package dto

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v3/object"
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type VkUserInfo struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Photo200  string `json:"photo_200"`
	Photo100  string `json:"photo_100"`
}

type User struct {
	Id           uint           `json:"id,omitempty"`
	Username     string         `json:"username,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt,omitempty"`
	PersonalInfo *VkUserInfo    `json:"personalInfo,omitempty"`
}

func VkUserInfoWrap(u *object.UsersUser) *VkUserInfo {
	if u == nil {
		return nil
	}

	return &VkUserInfo{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Photo100:  u.Photo100,
		Photo200:  u.Photo200,
	}
}

func UserWrap(u *models.User) *User {
	if u == nil {
		return nil
	}

	vkAllUserInfo, err := models.GetVkUserInfo(u.VkId)
	if err != nil {
		fmt.Println("User info not got")
	}

	var vkCurrentUserInfo *VkUserInfo
	if vkAllUserInfo != nil {
		vkCurrentUserInfo = VkUserInfoWrap(vkAllUserInfo)
	} else {
		vkCurrentUserInfo = nil
	}

	return &User{
		Id:           u.ID,
		Username:     u.Username,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		DeletedAt:    u.DeletedAt,
		PersonalInfo: vkCurrentUserInfo,
	}
}
