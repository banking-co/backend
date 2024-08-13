package base

import (
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
)

func StartApp(req *entities.Request) {
	var db = mysqldb.DB

	user, personalUserInfo, err := models.GetUserByUsername(db, req.StartParams.VkUserID)
	if err != nil {
		req.SendError(types.ErrorCodeBadRequest)
		return
	}

	req.SendMessage(req.Event, dto.ResponseStartApp{
		User:     dto.UserWrap(user, personalUserInfo),
		Bans:     dto.BansWrap(user.Bans),
		Balances: dto.BalancesWrap(user.Balances),
	})

	if len(user.Bans) >= 1 {
		req.SendError(types.ErrorCodeIsBanned)
		req.Disconnect()
		return
	}
}
