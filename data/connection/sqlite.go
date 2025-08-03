package connection

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func GetSqlite() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./valhalla.db")
	if err != nil {
		return nil, err
	}

	return db
}
