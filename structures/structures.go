package structures

import (
	"encoding/json"
	"rabotyaga-go-backend/types"
)

type EventParams struct {
	Event types.EventType `json:"event"`
	Data  json.RawMessage `json:"data,omitempty"`
}

type ResponseUserGet struct {
	Id        *uint   `json:"id,omitempty"`
	UserId    *uint   `json:"userId,omitempty"`
	Username  *string `json:"username,omitempty"`
	CreatedAt *uint   `json:"createdAt,omitempty"`
	UpdatedAt *uint   `json:"updatedAt,omitempty"`
	DeletedAt *uint   `json:"deletedAt,omitempty"`
}

type StartApp struct {
	User ResponseUserGet `json:"user,omitempty"`
}

type Error struct {
	Msg string `json:"msg"`
}
