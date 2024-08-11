package dto

import "rabotyaga-go-backend/models"

type BusinessStaff struct {
	Id uint `json:"id"`
}

func BusinessStaffWrap(b *models.BusinessStaff) *BusinessStaff {
	if b == nil {
		return nil
	}

	return &BusinessStaff{
		Id: b.ID,
	}
}
