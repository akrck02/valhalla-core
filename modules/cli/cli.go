package cli

import (
	"fmt"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/tables"

	"github.com/akrck02/valhalla-core/sdk/logger"
)

func Start() {

	// Connect to database
	db, err := database.Connect("./valhalla.db")
	if nil != err {
		logger.Errorf(err)
		return
	}

	defer db.Close()
	fmt.Println("Connected to the SQLite database successfully.")

	// Create tables
	err = tables.UpdateDatabaseTablesToLatestVersion(".", tables.MainDatabase, db)
	if nil != err {
		logger.Errorf(err)
		return
	}

	// Get the version of SQLite
	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sqliteVersion)
}
