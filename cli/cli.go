package cli

import (
	"fmt"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/tables"

	"github.com/akrck02/valhalla-core/logger"
)

func Start() {

	// Connect to database
	db, err := database.Connect()
	if nil != err {
		logger.Errorf(err)
		return
	}

	defer db.Close()
	fmt.Println("Connected to the SQLite database successfully.")

	// Create tables
	err = tables.CreateTables(db)
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
