package database

import (
	"zcelero/database/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Migrator().CreateTable(&entity.TextManagement{})
}
