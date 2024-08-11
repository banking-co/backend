package utils

import (
	"encoding/json"
	"rabotyaga-go-backend/types"
)

type SendData struct {
	Event types.EventType `json:"event"`
	Data  interface{}     `json:"data"`
}

func MarshalData(e types.EventType, d interface{}) ([]byte, error) {
	data, err := json.Marshal(SendData{Event: e, Data: d})
	if err != nil {
		return nil, err
	}

	return data, nil
}
