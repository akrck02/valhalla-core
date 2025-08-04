package database

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./valhalla.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
