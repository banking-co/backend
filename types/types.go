package types

type EventType = string
type ErrorMessage = string

const (
	EventPing               EventType = "ping"
	EventPong               EventType = "pong"
	EventStartApp           EventType = "start_app"
	EventGetBusiness        EventType = "get_business"
	EventGetPrimaryBusiness EventType = "get_pr_business"
	EventBalanceGet         EventType = "balance_get"
	EventUserGet            EventType = "user_get"
	EventError              EventType = "error"
)

const (
	BusinessRoleBot = 0
)
