package business

import (
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
)

func Get(req *entities.Request) {
	var db = mysqldb.DB
	var bu *models.Business
	var uID = req.PickInt("userId")
	var bID = req.PickInt("businessId")

	if uID == nil && bID == nil {
		req.SendError(types.ErrorCodeBadRequest)
		return
	}

	if bID != nil {
		b, err := models.GetBusinessById(db, *bID)
		if err != nil {
			req.SendError(types.ErrorCodeBadRequest)
			return
		}

		bu = b
	}

	if uID != nil {
		b, err := models.GetBusinessByUserId(db, *uID)
		if err != nil {
			req.SendError(types.ErrorCodeBadRequest)
			return
		}

		bu = b
	}

	user, personalUserInfo, err := models.GetUserById(db, bu.UserID)
	if err != nil {
		req.SendError(types.ErrorCodeInternalServerError)
		return
	}

	req.SendMessage(req.Event, dto.ResponseBusinessGet{
		BusinessID:    bu.ID,
		Business:      dto.BusinessWrap(bu),
		BusinessRoles: dto.BusinessRolesWrap(bu.Roles),
		User:          dto.UserWrap(user, personalUserInfo),
	})
}
