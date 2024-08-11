package business

import (
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
)

func GetStaff(req *entities.Request) {
	var db = mysqldb.DB
	var bID = req.PickInt("businessId")

	staff, err := models.GetBusinessStaffByBusinessId(db, bID)
	if err != nil {
		req.SendError(types.ErrorCodeBadRequest)
		return
	}

	req.SendMessage(req.Event, dto.ResponseBusinessStaffGet{
		BusinessID:    bID,
		BusinessStaff: dto.BusinessStaffsWrap(staff),
	})
}
