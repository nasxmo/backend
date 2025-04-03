package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate models
	// if err := db.AutoMigrate(&models.User{}); err != nil {
	//     return nil, fmt.Errorf("failed to migrate database: %w", err)
	// }

	return db, nil
}