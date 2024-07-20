package structures

import (
	"encoding/json"
	"rabotyaga-go-backend/types"
)

type EventParams struct {
	Event types.EventType `json:"event"`
	Data  json.RawMessage `json:"data,omitempty"`
}

type Error struct {
	Msg string `json:"msg"`
}

type User struct {
	Id        uint   `json:"id,omitempty"`
	UserId    uint   `json:"userId,omitempty"`
	Username  string `json:"username,omitempty"`
	CreatedAt uint   `json:"createdAt,omitempty"`
	UpdatedAt uint   `json:"updatedAt,omitempty"`
	DeletedAt uint   `json:"deletedAt,omitempty"`
}

type RequestUserGet struct {
	UserId uint `json:"userId,omitempty"`
}

type ResponseUserGet struct {
	User User `json:"user"`
}

type RequestBalance struct {
	UserId string `json:"userId"`
}

type ResponseBalance struct {
	BalanceUsd string `json:"balanceUsd"`
}

type RequestStartApp struct {
	Token string `json:"token"`
}

type ResponseStartApp struct {
	User    *User            `json:"user,omitempty"`
	Balance *ResponseBalance `json:"balance,omitempty"`
}
