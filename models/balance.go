package models

import (
	"errors"
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	UserID   uint   `gorm:"not null;index"`
	Amount   uint64 `gorm:"not null;default:0"`
	Currency string `gorm:"type:varchar(6);not null;default:'usd'"`
	User     User   `gorm:"foreignKey:UserID"`
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
