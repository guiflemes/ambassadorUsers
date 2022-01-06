package database

import (
	"users/src/models"
)

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
