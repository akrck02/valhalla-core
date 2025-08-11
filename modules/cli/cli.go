package cli

import (
	"fmt"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/tables"

	"github.com/akrck02/valhalla-core/sdk/logger"
)

func Start() {

	err := database.StartConnectionPool()
	if nil != err {
		logger.Errorf(err)
		return
	}

	// Connect to database
	db := database.GetConnection()
	defer database.ReturnConnection(db)
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
		logger.Errorf(err)
		return
	}

	fmt.Println(sqliteVersion)
}
