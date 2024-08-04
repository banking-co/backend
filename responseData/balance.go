package responseData

import (
	"rabotyaga-go-backend/models"
	"time"
)

type Balance struct {
	Id        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	Amount    uint64    `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func BalanceWrap(b *models.Balance) *Balance {
	if b == nil {
		return nil
	}

	return &Balance{
		Id:        b.ID,
		UserID:    b.UserID,
		Amount:    b.Amount,
		Currency:  b.Currency,
		CreatedAt: b.CreatedAt,
		UpdateAt:  b.UpdatedAt,
	}
}

func BalancesWrap(balances []models.Balance) []*Balance {
	var bA = make([]*Balance, 0, len(balances))

	for _, b := range balances {
		bCopy := b
		bA = append(bA, BalanceWrap(&bCopy))
	}

	return bA
}
