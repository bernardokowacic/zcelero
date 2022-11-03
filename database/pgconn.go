package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CreatePGConn opens the database connection
func CreatePGConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=database user=postgres password=%s dbname=posterr port=5432 sslmode=disable TimeZone=America/Sao_Paulo", os.Getenv("POSTGRES_PASSWORD"))
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return nil, err
	}

	return db, nil
}
