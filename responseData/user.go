package responseData

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v3/api"
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type VkUserInfo struct {
	Photo200 string `json:"photo_200"`











	 string `json:"photo_200"`
}

type User struct {
	Id        uint                  `json:"id,omitempty"`
	Username  string                `json:"username,omitempty"`
	CreatedAt time.Time             `json:"createdAt,omitempty"`
	UpdatedAt time.Time             `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt        `json:"deletedAt,omitempty"`
	VkData    *api.UsersGetResponse `json:"vkData,omitempty"`
}

func UserWrap(u *models.User) *User {
	if u == nil {
		return nil
	}

	vkAllUserInfo, err := models.GetVkUserInfo(u.VkId)
	if err != nil {
		fmt.Println("User info not got")
	}

	var vkCurrentUserInfo
	for _, vkU := range *vkAllUserInfo {
		if vkU.ID == u.VkId {
		}
	}


	return &User{
		Id:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
		VkData:    vkData,
	}
}
