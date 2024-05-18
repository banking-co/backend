package base

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
)

func Ping(conn net.Conn, code ws.OpCode, _ json.RawMessage) {
	err := wsutil.WriteServerMessage(conn, code, []byte("Pong"))
	if err != nil {
		return
	}
}
