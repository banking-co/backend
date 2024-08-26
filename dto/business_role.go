package dto

import (
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type BusinessRole struct {
	Id         uint `json:"id"`
	BusinessID uint `json:"bankId"`

	RoleId   uint8  `json:"roleId"`
	RoleName string `json:"roleName"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func BusinessRoleWrap(b *models.BusinessRole) *BusinessRole {
	if b == nil {
		return nil
	}

	return &BusinessRole{
		Id:         b.ID,
		BusinessID: b.BusinessID,

		RoleId:   b.RoleId,
		RoleName: b.RoleName,

		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
	}
}

func BusinessRolesWrap(bR []models.BusinessRole) []*BusinessRole {
	var bA = make([]*BusinessRole, 0, len(bR))

	for _, b := range bR {
		bCopy := b
		bA = append(bA, BusinessRoleWrap(&bCopy))
	}

	return bA
}
