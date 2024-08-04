package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"rabotyaga-go-backend/models"
)

type Options struct {
	Database       string `json:"Database"`
	Username       string `json:"Username"`
	Host           string `json:"Host"`
	MaxConnections string `json:"MaxConnections"`
	Port           string `json:"Port"`
	Password       string `json:"Password"`
}

var modelList = []interface{}{
	models.User{},
}

var DB *gorm.DB

func Init(opt Options) {
	dsn := opt.Username + ":" + opt.Password + "@tcp(" + opt.Host + ":" + opt.Port + ")/" + opt.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
}

func Migrate() {
	if err := models.RegisterModels(DB); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
