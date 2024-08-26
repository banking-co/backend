package dto

import (
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"time"
)

type BusinessStaff struct {
	ID         uint `json:"id"`
	BusinessID uint `json:"bankId"`

	UserType uint8 `json:"userType"`
	RoleID   uint8 `json:"roleId"`
	Salary   int   `json:"salary"`

	EmployerID uint `json:"employerId"`
	WorkerID   uint `json:"workerId"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func BusinessStaffWrap(b *models.BusinessStaff) *BusinessStaff {
	if b == nil {
		return nil
	}

	return &BusinessStaff{
		ID:         b.ID,
		BusinessID: b.BusinessID,

		UserType: b.UserType,
		RoleID:   b.RoleID,
		Salary:   b.Salary,

		EmployerID: b.EmployerID,
		WorkerID:   b.WorkerID,

		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
	}
}

func BusinessStaffsWrap(b []*models.BusinessStaff) []*BusinessStaff {
	if b == nil {
		return nil
	}

	staff := make([]*BusinessStaff, 0, len(b))

	for _, s := range b {
		staff = append(staff, BusinessStaffWrap(s))
	}

	return staff
}
