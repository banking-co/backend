package base

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
)

func StartApp(conn net.Conn, code ws.OpCode, data json.RawMessage) {
	_, err := utils.UnmarshalData[structures.RequestStartApp](data)
	user := structures.User{
		Id:        1,
		UserId:    1,
		CreatedAt: 312312321,
		DeletedAt: 12312312,
		UpdatedAt: 412312312412,
		Username:  "d_maximyuk",
	}

	if err == nil {
		resData, err := utils.MarshalData[structures.ResponseStartApp](types.EventStartApp, &structures.ResponseStartApp{User: &user})

		if err != nil {
			return
		}

		fmt.Println("send")
		err = wsutil.WriteServerMessage(conn, code, resData)
		if err != nil {
			return
		}
	}
}
