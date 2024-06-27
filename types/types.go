package types

type EventType = string
type ErrorMessage = string

const (
	EventPing       EventType = "ping"
	EventPong       EventType = "pong"
	EventStartApp   EventType = "start_app"
	EventBalanceGet EventType = "balance_get"
	EventUserGet    EventType = "user_get"
	EventError      EventType = "error"
)

const (
	//ErrorMessageNil          ErrorMessage = "data is nil"
	ErrorMessageParseData ErrorMessage = "data is nil"
	ErrorMessageMsgLength ErrorMessage = "data is nil"
	//ErrorMessageUidUndefined ErrorMessage = "the user ID is not defined or a conversion error has occurred"
	ErrorMessageMissingEvent ErrorMessage = "event is missing"
	ErrorMessageMissingData  ErrorMessage = "data is missing"
)
