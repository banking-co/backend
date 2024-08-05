package models

import "gorm.io/gorm"

func RegisterModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Business{},
		&Ban{},
		&Bonus{},
		&Balance{},
		&BusinessRole{},
		&BusinessStaff{},
		&BusinessUpgrade{},
		&BusinessProfit{},
		&Inventory{},
		&Auction{},
		&Item{},
		&Referral{},
		&PromoCode{},
		&Transaction{},
	)
}
