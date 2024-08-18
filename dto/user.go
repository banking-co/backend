package dto

import (
	"github.com/SevereCloud/vksdk/v3/object"
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type VkUserInfo struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	PhotoMax  string `json:"photoMax"`
	Photo200  string `json:"photo200"`
	Photo100  string `json:"photo100"`
	Photo50   string `json:"photo50"`
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
		Photo50:   u.Photo50,
		Photo100:  u.Photo100,
		Photo200:  u.Photo200,
		PhotoMax:  u.PhotoMax,
	}
}

func UserWrap(u *models.User, p *object.UsersUser) *User {
	if u == nil {
		return nil
	}

	return &User{
		Id:           u.ID,
		Username:     u.Username,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		DeletedAt:    u.DeletedAt,
		PersonalInfo: VkUserInfoWrap(p),
	}
}

func UsersWrap(us []*models.User, pIs map[int]*object.UsersUser) []*User {
	var users = make([]*User, 0, len(us))
	for _, u := range us {
		if u != nil {
			personalInfo := pIs[u.VkId]
			users = append(users, UserWrap(u, personalInfo))
		}
	}

	return users
}
