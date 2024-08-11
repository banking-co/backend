package types

import "encoding/json"

type EventType = string
type ErrorMessage = string

const (
	EventPing               EventType = "ping"
	EventPong               EventType = "pong"
	EventStartApp           EventType = "start_app"
	EventGetBusiness        EventType = "get_business"
	EventGetPrimaryBusiness EventType = "get_pr_business"
	EventGetBusinessStaff   EventType = "get_st_business"
	EventBalanceGet         EventType = "balance_get"
	EventBonusGet           EventType = "bonus_get"
	EventBonusReceive       EventType = "bonus_receive"
	EventUserGet            EventType = "user_get"
	EventError              EventType = "error"
)

var Events = []EventType{
	EventPing,
	EventPong,
	EventStartApp,
	EventGetPrimaryBusiness,
	EventGetBusinessStaff,
	EventGetBusiness,
	EventBalanceGet,
	EventBonusGet,
	EventBonusReceive,
	EventUserGet,
	EventError,
}

type EventParams struct {
	Event EventType       `json:"event"`
	Data  json.RawMessage `json:"data,omitempty"`
}

type Error struct {
	Msg string `json:"msg"`
}
