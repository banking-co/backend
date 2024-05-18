package utils

import "encoding/json"

func MarshalData(d any) ([]byte, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return data, nil
}
