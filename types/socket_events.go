package types

import "encoding/json"

type EventType = string
type ErrorCode = string

const (
	ErrorCodeIsBanned            ErrorCode = "is_banned"
	ErrorCodeBadRequest          ErrorCode = "bad_request"
	ErrorCodeInternalServerError ErrorCode = "internal_server_error"
	ErrorCodeForbidden           ErrorCode = "forbidden"
)

const (
	EventPing     EventType = "ping"
	EventPong     EventType = "pong"
	EventStartApp EventType = "start_app"
	EventError    EventType = "error"

	EventUserGet EventType = "get_usr"

	EventGetBusiness             EventType = "get_bus"
	EventGetBusinessStaff        EventType = "get_st_bus"
	EventGetBusinessStaffRecruit EventType = "get_st_r_bus"
	EventBuyBusinessStaffRecruit EventType = "buy_st_r_bus"

	EventBalanceGet EventType = "bal_get"

	EventBonusGet     EventType = "bon_get"
	EventBonusReceive EventType = "bon_receive"
)

type EventParams struct {
	Event EventType       `json:"event"`
	Data  json.RawMessage `json:"data,omitempty"`
}
