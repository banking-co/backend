package models

import (
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);unique;not null"`
}
