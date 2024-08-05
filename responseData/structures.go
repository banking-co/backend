package responseData

import (
	"encoding/json"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/types"
)

type EventParams struct {
	Event types.EventType `json:"event"`
	Data  json.RawMessage `json:"data,omitempty"`
}

type Error struct {
	Msg string `json:"msg"`
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

type ResponseError struct {
	Message string `json:"message,omitempty"`
}

type ResponseStartApp struct {
	User     *User      `json:"user,omitempty"`
	Balances []*Balance `json:"balances,omitempty"`
	Bans     []*Ban     `json:"bans,omitempty"`
}

type ResponseBusinessGet struct {
	Business        *Business               `json:"bank,omitempty"`
	BusinessStaff   *models.BusinessStaff   `json:"bankStaff,omitempty"`
	BusinessRole    *models.BusinessRole    `json:"bankRoles,omitempty"`
	BusinessUpgrade *models.BusinessUpgrade `json:"bankUpgrades,omitempty"`
	BusinessProfits *models.BusinessProfit  `json:"bankProfits,omitempty"`
}
