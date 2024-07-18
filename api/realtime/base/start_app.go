package base

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/utils"
)

func StartApp(conn net.Conn, code ws.OpCode, data json.RawMessage) {
	reqData, err := utils.UnmarshalData[interface{}](data)

	if err == nil {
		resData, err := utils.MarshalData(reqData)
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
