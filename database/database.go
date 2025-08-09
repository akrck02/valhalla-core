package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func Connect(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", fmt.Sprintf("%s?cache=shared&mode=rwc&_journal_mode=WAL", filePath))
	if err != nil {
		return nil, err
	}

	return db, nil
}
