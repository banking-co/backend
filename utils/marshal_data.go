package utils

import (
	"encoding/json"
	"rabotyaga-go-backend/types"
)

type SendData struct {
	Event types.EventType `json:"event"`
	Data  interface{}     `json:"data"`
}

/*
	func UnmarshalData[T any](bytes []byte) (*T, error) {
		out := new(T)
		if err := json.Unmarshal(bytes, out); err != nil {
			return nil, err
		}
		return out, nil
	}
*/

func MarshalData[T any](e types.EventType, d *T) ([]byte, error) {
	data, err := json.Marshal(SendData{Data: d, Event: e})
	if err != nil {
		return nil, err
	}

	return data, nil
}
