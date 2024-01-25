package database

import (
	"github.com/Ajay1290/go-auth/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:admin@/godb"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database.")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})

}
