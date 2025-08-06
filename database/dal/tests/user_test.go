package tests

import (
	"fmt"
	"testing"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/dal"
	"github.com/akrck02/valhalla-core/database/tables"
	"github.com/akrck02/valhalla-core/sdk/logger"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func TestRegister(t *testing.T) {

	basePath := "../../../"
	filePath := "../../../valhalla-test.db"
	db, _ := database.Connect(filePath)
	defer db.Close()

	erri := tables.UpdateDatabaseTablesToLatestVersion(basePath, db)
	if erri != nil {
		t.Fail()
		logger.Error(erri.Error())
	}

	_, err := dal.RegisterUser(db, &models.User{Email: "akrck02@gmail.com"}, &models.UserAuth{Password: "$P4ssw0rdW3db1"})

	if err != nil {
		t.Fail()
		logger.Error(err)
	}

	fmt.Print("")
}
