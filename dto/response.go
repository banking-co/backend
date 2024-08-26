package dto

import (
	"time"
)

type ResponseError struct {
	Code string `json:"code,omitempty"`
}

type ResponsePingGet struct {
	Time time.Time `json:"time"`
}

type ResponseStartApp struct {
	User     *User      `json:"user,omitempty"`
	Balances []*Balance `json:"balances,omitempty"`
	Bans     []*Ban     `json:"bans,omitempty"`
}

type ResponseUserGet struct {
	User          *User           `json:"user,omitempty"`
	Work          any             `json:"work,omitempty"`
	Business      *Business       `json:"bank,omitempty"`
	BusinessRoles []*BusinessRole `json:"bankRoles,omitempty"`
}

type ResponseBusinessGet struct {
	Type          string          `json:"type"`
	BusinessID    uint            `json:"bankId,omitempty"`
	Business      *Business       `json:"bank,omitempty"`
	BusinessRoles []*BusinessRole `json:"bankRoles,omitempty"`
	User          *User           `json:"user,omitempty"`
}

type ResponseBusinessStaffGet struct {
	BusinessID    *int             `json:"bankId,omitempty"`
	BusinessStaff []*BusinessStaff `json:"bankStaff,omitempty"`
	Users         []*User          `json:"users,omitempty"`
}

type ResponseBusinessStaffRecruitGet struct {
	Items []*Item `json:"items,omitempty"`
}
