package responseData

import (
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type User struct {
	Id        uint           `json:"id,omitempty"`
	Username  string         `json:"username,omitempty"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

func UserWrap(u models.User) User {
	return User{
		Id:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
