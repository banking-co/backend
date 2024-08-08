package responseData

import (
	"fmt"
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type VkUserInfo struct {
	Photo200  string `json:"photo_200"`
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type User struct {
	Id        uint           `json:"id,omitempty"`
	Username  string         `json:"username,omitempty"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
	VkData    VkUserInfo     `json:"vkData,omitempty"`
}

func UserWrap(u *models.User) *User {
	if u == nil {
		return nil
	}

	vkAllUserInfo, err := models.GetVkUserInfo(u.VkId)
	if err != nil {
		fmt.Println("User info not got")
	}

	var vkCurrentUserInfo VkUserInfo
	for _, vkU := range *vkAllUserInfo {
		if vkU.ID == u.VkId {
			vkCurrentUserInfo = VkUserInfo{
				Id:        vkU.ID,
				FirstName: vkU.FirstName,
				LastName:  vkU.LastName,
				Photo200:  vkU.Photo200,
			}
		}
	}

	return &User{
		Id:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
		VkData:    vkCurrentUserInfo,
	}
}
