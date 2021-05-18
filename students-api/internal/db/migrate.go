package db

import (
	"students-api/internal/service/student"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&student.Student{}); err != nil {
		return err
	}

	return nil
}
