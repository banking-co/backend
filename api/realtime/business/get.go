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
	var uID = req.PickUint("userId")
	var bID = req.PickUint("businessId")
	var bType = req.PickString("type")

	if bType == nil {
		newBusinessType := "default"
		bType = &newBusinessType
	}

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
		Type:          *bType,
		BusinessID:    bu.ID,
		Business:      dto.BusinessWrap(bu),
		BusinessRoles: dto.BusinessRolesWrap(bu.Roles),
		User:          dto.UserWrap(user, personalUserInfo),
	})
}
