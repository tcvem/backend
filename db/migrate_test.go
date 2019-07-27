package db

import (
	"database/sql"

	"testing"
)

func TestMigrateDB(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres password=postgres sslmode=disable dbname=backend")
	if err != nil {
		t.Fatalf("unable to connect to backend database: %v", err)
	}
	err = MigrateDB(db)
	if err != nil {
		t.Fatalf("unable to create backend database: %v", err)
	}
}
