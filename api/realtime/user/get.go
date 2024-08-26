package user

import (
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
)

func Get(req *entities.Request) {
	db := mysqldb.DB
	uid := req.PickUint("uid")
	if uid == nil {
		req.SendError(types.ErrorCodeBadRequest)
		return
	}

	user, personalInfo, err := models.GetUserById(db, *uid)
	if err != nil {
		req.SendError(types.ErrorCodeInternalServerError)
		return
	}

	business, err := models.GetBusinessByUserId(db, user.ID)
	if err != nil {
		req.SendError(types.ErrorCodeInternalServerError)
		return
	}

	businessRoles, err := models.GetBusinessRolesByBusinessId(db, business.ID)
	if err != nil {
		req.SendError(types.ErrorCodeInternalServerError)
		return
	}

	userWork, err := models.GetUserWork(db, user.ID)
	if err != nil {
		req.SendError(types.ErrorCodeInternalServerError)
		return
	}

	req.SendMessage(req.Event, dto.ResponseUserGet{
		User:          dto.UserWrap(user, personalInfo),
		Business:      dto.BusinessWrap(business),
		BusinessRoles: dto.BusinessRolesWrap(businessRoles),
		Work:          dto.BusinessStaffWrap(userWork),
	})
}
