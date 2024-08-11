package utils

import "rabotyaga-go-backend/types"

func GetBigLenEvent() uint8 {
	var maxEventLength uint8

	for _, e := range types.Events {
		if uint8(len(e)) > maxEventLength {
			maxEventLength = uint8(len(e))
		}
	}

	return maxEventLength
}
