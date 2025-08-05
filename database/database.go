package database

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func Connect(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filePath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
