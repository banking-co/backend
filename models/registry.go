package models

import "gorm.io/gorm"

func RegisterModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Ban{},
		&Bonus{},
		&Business{},
		&BusinessUpgrade{},
		&BusinessStaff{},
		&BusinessProfit{},
		&Inventory{},
		&Auction{},
		&Item{},
		&Referral{},
		&PromoCode{},
		&Transaction{},
	)
}
