package balance

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
)

func Get(e types.EventType, conn net.Conn, code ws.OpCode, vkParams *vkapps.Params, data json.RawMessage) {
	var db = mysqldb.DB

	balances, err := models.GetBalancesByUid(db, vkParams.VkUserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	resData, err := utils.MarshalData[dto.ResponseBalancesGet](e, &dto.ResponseBalancesGet{
		Balances: dto.BalancesWrap(balances),
	})
	if err != nil {
		return
	}

	err = wsutil.WriteServerMessage(conn, code, resData)
	if err != nil {
		return
	}
}
