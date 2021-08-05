package db

import (
	"auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:pass@/mydb"), &gorm.Config{})
	if err != nil {
		panic("Couldnt connect to the database")
	}
	conn.AutoMigrate(&models.User{})
}
