package base

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/types"
)

func Ping(conn net.Conn, code ws.OpCode, _ *vkapps.Params, _ json.RawMessage) {
	err := wsutil.WriteServerMessage(conn, code, []byte(types.EventPong))
	if err != nil {
		return
	}
}
