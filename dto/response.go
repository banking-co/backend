package dto

import "rabotyaga-go-backend/models"

type ResponseError struct {
	Code string `json:"code,omitempty"`
}

type ResponseStartApp struct {
	User     *User      `json:"user,omitempty"`
	Balances []*Balance `json:"balances,omitempty"`
	Bans     []*Ban     `json:"bans,omitempty"`
}

type ResponseUserGet struct {
	User User `json:"user"`
}

type ResponseBusinessGet struct {
	Business        *Business               `json:"bank,omitempty"`
	BusinessStaff   *models.BusinessStaff   `json:"bankStaff,omitempty"`
	BusinessRole    *models.BusinessRole    `json:"bankRoles,omitempty"`
	BusinessUpgrade *models.BusinessUpgrade `json:"bankUpgrades,omitempty"`
	BusinessProfits *models.BusinessProfit  `json:"bankProfits,omitempty"`
}

type ResponseBusinessStaffGet struct {
	BusinessID    *int                    `json:"bankId,omitempty"`
	BusinessStaff *[]models.BusinessStaff `json:"bankStaff,omitempty"`
}
