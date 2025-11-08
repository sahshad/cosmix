package database

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

func AutoMigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	)
}
