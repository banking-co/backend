package dto

type RequestUserGet struct {
	UserId uint `json:"userId,omitempty"`
}

type RequestBalance struct {
	UserId string `json:"userId"`
}

type RequestBusinessGet struct {
	UserId     *int `json:"userId,omitempty"`
	BusinessId *int `json:"businessId,omitempty"`
}

type ResponseBalance struct {
	BalanceUsd string `json:"balanceUsd"`
}

type RequestStartApp struct {
	Token string `json:"token"`
}
