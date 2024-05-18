package user

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/mysql"
	"rabotyaga-go-backend/utils"
)

type RequestUserGet struct {
	userId uint
}

func Get(conn net.Conn, code ws.OpCode, data json.RawMessage) {
	reqData, err := utils.UnmarshalData[RequestUserGet](data)

	if err == nil {
		user, _ := mysql.USER_GET_BY_UID(reqData.userId)

		resData, err := utils.MarshalData(user)
		if err != nil {
			return
		}

		err = wsutil.WriteServerMessage(conn, code, resData)
		if err != nil {
			return
		}
	}
}
