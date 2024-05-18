package balance

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
)

func Get(conn net.Conn, code ws.OpCode, data json.RawMessage) {
	err := wsutil.WriteClientMessage(conn, code, data)
	if err != nil {
		return
	}
}
