package models

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type Balance struct {
	gorm.Model
	UserID     uint   `gorm:"not null;index"`
	Importance uint8  `gorm:"not null"`
	Amount     uint64 `gorm:"not null;default:0"`
	Currency   string `gorm:"type:varchar(6);not null;default:'usd'"`
	User       User   `gorm:"foreignKey:UserID"`
}

func (b *Balance) AdjustAmount(tx *gorm.DB, delta int64) error {
	newAmount := int64(b.Amount) + delta
	if newAmount < 0 {
		return errors.New("balance cannot be negative")
	}

	b.Amount = uint64(newAmount)
	return tx.Save(b).Error
}

func (b *Balance) BeforeSave(tx *gorm.DB) (err error) {
	if b.Amount < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

func GetBalancesByUid(db *gorm.DB, uid int) ([]Balance, error) {
	var balances []Balance

	if err := db.Joins("JOIN users ON users.id = balances.user_id").
		Where("users.username = ?", "id"+strconv.Itoa(uid)).
		Find(&balances).Error; err != nil {
		return nil, err
	}

	return balances, nil
}
