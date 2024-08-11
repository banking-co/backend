package balance

import (
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
)

func Get(req *entities.Request) {
	var db = mysqldb.DB

	balances, err := models.GetBalancesByUid(db, req.StartParams.VkUserID)
	if err != nil {
		req.SendError(types.ErrorCodeBadRequest)
		return
	}

	req.SendMessage(req.Event, dto.ResponseBalancesGet{
		Balances: dto.BalancesWrap(balances),
	})
}
