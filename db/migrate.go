package db

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tcvem/backend/pkg/pb"
)

// MigrateDB builds the backend application database tables
func MigrateDB(dbSQL *sql.DB) error {
	db, err := gorm.Open("postgres", dbSQL)
	if err != nil {
		return err
	}
	defer db.Close()

	db.LogMode(true)

	// NOTE: Using db.AutoMigrate is a temporary measure to structure the contacts
	// database schema. The atlas-app-toolkit team will come up with a better
	// solution that uses database migration files.
	return db.AutoMigrate(
		&pb.CertficateORM{},
	).Error
}
