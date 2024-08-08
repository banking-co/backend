package dto

import (
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type Business struct {
	Id        uint           `json:"id"`
	UserID    uint           `json:"userId"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func BusinessWrap(b *models.Business) *Business {
	if b == nil {
		return nil
	}

	return &Business{
		Id:        b.ID,
		UserID:    b.UserID,
		Name:      b.Name,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
	}
}

func BusinessesWrap(bans []models.Business) []*Business {
	var bA = make([]*Business, 0, len(bans))

	for _, b := range bans {
		bCopy := b
		bA = append(bA, BusinessWrap(&bCopy))
	}

	return bA
}
