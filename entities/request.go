package entities

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
)

type Request struct {
	Event       types.EventType
	Data        json.RawMessage
	StartParams *vkapps.Params

	Conn net.Conn
	Op   ws.OpCode
}

func (r *Request) SendMessage(e types.EventType, d interface{}) {
	data, err := utils.MarshalData(e, &d)
	if err != nil {
		return
	}

	err = wsutil.WriteServerMessage(r.Conn, r.Op, data)
	if err != nil {
		fmt.Println("[SendMessage] Sending message ended with an error:", err)
		return
	}
}

func (r *Request) SendError(code types.ErrorCode) {
	r.SendMessage(types.EventError, dto.ResponseError{Code: code})
}

func (r *Request) Disconnect() {
	err := r.Conn.Close()

	if err != nil {
		fmt.Println("[Disconnect] Disconnect user ended with an error:", err)
	}
}

func (r *Request) PickInt(k string) *int {
	var data map[string]interface{}

	err := json.Unmarshal(r.Data, &data)
	if err != nil {
		fmt.Println("[Error] Failed to unmarshal JSON:", err)
		return nil
	}

	if val, ok := data[k]; ok {
		if floatVal, ok := val.(float64); ok {
			intVal := int(floatVal)
			return &intVal
		}
	}

	return nil
}

func (r *Request) PickUint(k string) *uint {
	var data map[string]interface{}

	err := json.Unmarshal(r.Data, &data)
	if err != nil {
		fmt.Println("[Error] Failed to unmarshal JSON:", err)
		return nil
	}

	if val, ok := data[k]; ok {
		if floatVal, ok := val.(float64); ok {
			intVal := uint(floatVal)
			return &intVal
		}
	}

	return nil
}

func (r *Request) PickString(k string) *string {
	var data map[string]interface{}

	err := json.Unmarshal(r.Data, &data)
	if err != nil {
		fmt.Println("[Error] Failed to unmarshal JSON:", err)
		return nil
	}

	if val, ok := data[k]; ok {
		if str, ok := val.(string); ok {
			return &str
		}
	}

	return nil
}
