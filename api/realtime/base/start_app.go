package base

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

func StartApp(e types.EventType, conn net.Conn, code ws.OpCode, vkParams *vkapps.Params, data json.RawMessage) {
	var db = mysqldb.DB

	user, err := models.GetUserByUsername(db, vkParams.VkUserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	resData, err := utils.MarshalData[dto.ResponseStartApp](e, &dto.ResponseStartApp{
		User:     dto.UserWrap(user),
		Bans:     dto.BansWrap(user.Bans),
		Balances: dto.BalancesWrap(user.Balances),
	})
	if err != nil {
		return
	}

	err = wsutil.WriteServerMessage(conn, code, resData)
	if err != nil {
		return
	}

	if len(user.Bans) >= 1 {
		err := conn.Close()
		if err != nil {
			return
		}
		return
	}
}
