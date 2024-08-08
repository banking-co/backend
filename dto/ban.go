package dto

import (
	"rabotyaga-go-backend/models"
	"time"
)

type Ban struct {
	Id          uint       `json:"id"`
	UserID      uint       `json:"userId"`
	InitiatorID uint       `json:"initiatorId"`
	Reason      string     `json:"reason"`
	CreatedAt   time.Time  `json:"createdAt"`
	Until       *time.Time `json:"until,omitempty"`
}

func BanWrap(b *models.Ban) *Ban {
	if b == nil {
		return nil
	}

	return &Ban{
		Id:          b.ID,
		UserID:      b.UserID,
		InitiatorID: b.InitiatorID,
		Reason:      b.Reason,
		CreatedAt:   b.CreatedAt,
		Until:       b.Until,
	}
}

func BansWrap(bans []models.Ban) []*Ban {
	var bA = make([]*Ban, 0, len(bans))

	for _, b := range bans {
		bCopy := b
		bA = append(bA, BanWrap(&bCopy))
	}

	return bA
}
