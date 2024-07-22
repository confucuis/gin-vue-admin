package utils

import (
	"fmt"
	"gin-vue-admin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	username := GetEnv("DB_USER")
	password := GetEnv("DB_PASS")
	host := GetEnv("DB_HOST")
	port := GetEnv("DB_PORT")
	dbName := GetEnv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{}, &models.Menu{})

	Db = db
}
