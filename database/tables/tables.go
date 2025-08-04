package tables

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func CreateTables(db *sql.DB) error {

	currentDatabaseVersion := 1

	// if no database exists, create one
	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='database_metadata'").Scan()
	if err == sql.ErrNoRows {
		return createTablesForVersion(db, 0, currentDatabaseVersion)
	}

	// get current version and update existing database
	var databaseVersion int
	err = db.QueryRow("SELECT version FROM database_metadata").Scan(&databaseVersion)
	if err != nil {
		return err
	}

	return createTablesForVersion(db, databaseVersion, currentDatabaseVersion)
}

func createTablesForVersion(db *sql.DB, currentVersion int, targetVersion int) error {

	for version := currentVersion + 1; version <= targetVersion; version++ {
		err := executeScriptIfExists(db, fmt.Sprintf("sql/v%d/tables.sql", version))
		if nil != err {
			return err
		}

		err = executeScriptIfExists(db, fmt.Sprintf("sql/v%d/data.sql", version))
		if nil != err {
			return err
		}
	}

	return nil
}

func executeScriptIfExists(db *sql.DB, fileName string) error {

	// if the file does not exist do not execute anything
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil
	}

	script, err := os.ReadFile(fileName)
	if nil != err {
		return err
	}

	statements := strings.Split(string(script), ";")

	for _, statement := range statements {
		_, err = db.Exec(statement)
		if nil != err {
			return err
		}
	}

	return nil
}
