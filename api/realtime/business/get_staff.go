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

	var usersId = make([]uint, 0, len(staff)*2)
	if staff != nil {
		for _, u := range staff {
			usersId = append(usersId, u.WorkerID)
			usersId = append(usersId, u.EmployerID)
		}
	}

	users, personalUsersInfo, err := models.GetUsersByIds(db, usersId)
	if err != nil {
		req.SendError(types.ErrorCodeInternalServerError)
		return
	}

	req.SendMessage(req.Event, dto.ResponseBusinessStaffGet{
		BusinessID:    bID,
		BusinessStaff: dto.BusinessStaffsWrap(staff),
		Users:         dto.UsersWrap(users, personalUsersInfo),
	})
}
