package database

import "server/models"

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.UserToken{},
	)
}
