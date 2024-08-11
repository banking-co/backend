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

	req.SendMessage(req.Event, dto.ResponseBusinessGet{
		Business: dto.BusinessWrap(bu),
	})
}
