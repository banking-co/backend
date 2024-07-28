package user

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/mysql"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
)

func Get(conn net.Conn, code ws.OpCode, _ *vkapps.Params, data json.RawMessage) {
	reqData, err := utils.UnmarshalData[structures.RequestUserGet](data)

	if err == nil {
		user, _ := mysql.USER_GET_BY_UID(reqData.UserId)

		resData, err := utils.MarshalData[structures.ResponseUserGet](types.EventUserGet, &structures.ResponseUserGet{User: *user})
		if err != nil {
			return
		}

		err = wsutil.WriteServerMessage(conn, code, resData)
		if err != nil {
			return
		}
	}
}
