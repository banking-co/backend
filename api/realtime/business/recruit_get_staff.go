package business

import (
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
)

func RecruitGetStaff(req *entities.Request) {
	var db = mysqldb.DB

	items, err := models.GetItemsByType(db, types.ItemTypeBusinessStaff)
	if err != nil {
		req.SendError(types.ErrorCodeBadRequest)
		return
	}

	req.SendMessage(req.Event, dto.ResponseBusinessStaffRecruitGet{
		Items: dto.ItemsWrap(items),
	})
}
