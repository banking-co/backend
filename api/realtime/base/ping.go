package base

import (
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/types"
)

func Ping(req *entities.Request) {
	req.SendMessage(types.EventPong, struct{}{})
}
