package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"rabotyaga-go-backend/models"
	"reflect"
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

func AutoMigrateAllModels(db *gorm.DB) error {
	for _, model := range modelList {
		modelType := reflect.TypeOf(model)
		if modelType.Kind() == reflect.Ptr {
			modelType = modelType.Elem()
		}
		fmt.Printf("Migrating model: %s\n", modelType.Name())
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate model %s: %w", modelType.Name(), err)
		}
	}
	return nil
}

func Migrate() {
	if err := AutoMigrateAllModels(DB); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
