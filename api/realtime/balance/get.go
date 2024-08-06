package balance

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/types"
)

func Get(e types.EventType, conn net.Conn, code ws.OpCode, _ *vkapps.Params, data json.RawMessage) {
	err := wsutil.WriteClientMessage(conn, code, data)
	if err != nil {
		return
	}
}
