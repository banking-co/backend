package base

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"net"
)

func StartApp(conn net.Conn, code ws.OpCode, data json.RawMessage) {
}
