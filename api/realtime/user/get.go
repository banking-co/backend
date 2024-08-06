package user

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"net"
	"rabotyaga-go-backend/types"
)

func Get(e types.EventType, conn net.Conn, code ws.OpCode, startParams *vkapps.Params, data json.RawMessage) {
	//var user models.User
	//var db = database.DB
	return
}
