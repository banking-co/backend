package mysqldb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"rabotyaga-go-backend/models"
)

type Options struct {
	Database       string
	Username       string
	Host           string
	MaxConnections string
	Port           string
	Password       string
}

var DB *gorm.DB

func Init() {
	dbPassword, dbPasswordExist := os.LookupEnv("DB_PASSWORD")
	if !dbPasswordExist {
		panic("DB_PASSWORD not set in environment")
	}

	opt := Options{
		Database:       "banking",
		Username:       "backend",
		Host:           "localhost",
		MaxConnections: "10",
		Port:           "3306",
		Password:       dbPassword,
	}

	dsn := opt.Username + ":" + opt.Password + "@tcp(" + opt.Host + ":" + opt.Port + ")/" + opt.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	Migrate()
}

func Migrate() {
	if err := models.RegisterModels(DB); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
