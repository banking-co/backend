package base

import (
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/types"
	"time"
)

func Ping(req *entities.Request) {
	req.SendMessage(types.EventPong, dto.ResponsePingGet{Time: time.Now()})
}
