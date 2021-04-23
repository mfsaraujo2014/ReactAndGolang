package database

import (
	"modulo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("auth:auth@/yt_go_auth"), &gorm.Config{})
	if err != nil {
		panic("não foi possivel conectar ao banco")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
