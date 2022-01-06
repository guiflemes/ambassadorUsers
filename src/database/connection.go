package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(users_db:3306)/users"), &gorm.Config{})

	if err != nil {
		log.Panic("Could not conect with the database!")
	}

}
