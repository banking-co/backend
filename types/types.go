package types

type EventType = string
type ErrorMessage = string

const (
	RequestPing       EventType = "ping"
	RequestStartApp   EventType = "start_app"
	RequestBalanceGet EventType = "balance_get"
	RequestUserGet    EventType = "user_get"
)

const (
	ResponsePong       EventType = "pong"
	ResponseStartApp   EventType = "start_app"
	ResponseBalanceGet EventType = "balance_get"
	ResponseUserGet    EventType = "user_get"
	ResponseError      EventType = "error"
)

const (
	//ErrorMessageNil          ErrorMessage = "data is nil"
	ErrorMessageParseData ErrorMessage = "data is nil"
	ErrorMessageMsgLength ErrorMessage = "data is nil"
	//ErrorMessageUidUndefined ErrorMessage = "the user ID is not defined or a conversion error has occurred"
	ErrorMessageMissingEvent ErrorMessage = "event is missing"
	ErrorMessageMissingData  ErrorMessage = "data is missing"
)
