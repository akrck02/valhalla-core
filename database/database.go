package database

import (
	"database/sql"
	"fmt"

	"github.com/akrck02/valhalla-core/sdk/logger"
	_ "github.com/mattn/go-sqlite3"
)

var max_conns = 2
var conns = make(chan *sql.DB, max_conns)

func StartConnectionPool() error {

	// for i := 0; i < max_conns; i++ {
	// 	conn, err := connect("valhalla.db")
	// 	if nil != err {
	// 		return err
	// 	}

	// 	conns <- conn
	// }

	return nil
}

func GetConnection() *sql.DB {
	db, err := Connect("valhalla.db")
	if nil != err {
		logger.Error(err)
	}

	db.SetMaxOpenConns(100)

	return db
}

func ReturnConnection(c *sql.DB) {
	if nil == c {
		return
	}
	c.Close()
	// conns <- c
}

func Connect(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?cache=shared&mode=rwc&_journal_mode=WAL", filePath))
	if err != nil {
		return nil, err
	}

	return db, nil
}
