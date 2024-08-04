package models

import (
	"gorm.io/gorm"
	"time"
)

type BusinessStaff struct {
	gorm.Model
	UserType    int        `gorm:"not null"` // bot - 0, user - 1
	BusinessID  uint       `gorm:"not null"`
	EmployerID  uint       `gorm:"not null"`
	WorkerID    uint       `gorm:"not null"` // bot - 0, user - user_id
	Role        int        `gorm:"not null"` // the position that the person holds
	Salary      int        `gorm:"not null"` // per hours
	JoinedAt    time.Time  `gorm:"autoCreateTime"`
	DismissedAt *time.Time `gorm:"default:null"`
	Business    Business   `gorm:"foreignKey:BusinessID"`
	Employer    User       `gorm:"foreignKey:EmployerID"`
	Worker      User       `gorm:"foreignKey:WorkerID"`
}
