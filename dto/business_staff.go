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

	//var uIDs = make([]int, 0, 2)
	//var wU *models.User
	//var eU *models.User
	//if b.WorkerID != 0 {
	//	u, err := models.GetUserById(mysqldb.DB, b.WorkerID)
	//	if err != nil {
	//		return nil
	//	}
	//
	//	wU = u
	//	uIDs = append(uIDs, u.VkId)
	//}
	//
	//if b.EmployerID != 0 {
	//	u, err := models.GetUserById(mysqldb.DB, b.EmployerID)
	//	if err != nil {
	//		return nil
	//	}
	//
	//	eU = u
	//	uIDs = append(uIDs, u.VkId)
	//}
	//
	//var pU *[]object.UsersUser
	//if len(uIDs) >= 1 {
	//	i, err := models.GetVkUsersInfo(uIDs)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	pU = i
	//}
	//
	//var wPI *object.UsersUser
	//var ePI *object.UsersUser
	//if len(*pU) >= 1 {
	//	for _, u := range *pU {
	//		if wU != nil && u.ID == wU.VkId {
	//			wPI = &u
	//		}
	//
	//		if eU != nil && u.ID == eU.VkId {
	//			ePI = &u
	//		}
	//	}
	//}

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
