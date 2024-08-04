package base

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/database"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/responseData"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
)

func StartApp(conn net.Conn, code ws.OpCode, vkParams *vkapps.Params, data json.RawMessage) {
	var db = database.DB

	user, err := models.GetUserByUsername(db, vkParams.VkUserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	resData, err := utils.MarshalData[responseData.ResponseStartApp](types.EventStartApp, &responseData.ResponseStartApp{
		User:     responseData.UserWrap(user),
		Bans:     responseData.BansWrap(user.Bans),
		Balances: responseData.BalancesWrap(user.Balances),
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
