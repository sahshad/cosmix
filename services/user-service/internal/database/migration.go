package database

import (
	"user-service/internal/models"

	"gorm.io/gorm"
)

func AutoMigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.UserProfile{},
		&models.Follow{},
	)
}
