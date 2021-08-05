package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:pass@/mydb"), &gorm.Config{})
	if err != nil {
		panic("Couldnt connect to the database")
	}
}
