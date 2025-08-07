package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/tables"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/logger"
)

const TestDatabasePath string = "../../.."

func Assert(t *testing.T, predicate bool, failMessage string) {
	if !predicate {
		logger.Log("Test failed:", failMessage)
		t.FailNow()
	}
}

func AssertVErrorDoesNotExist(t *testing.T, error *errors.VError) {
	if nil != error {
		logger.Error("Test failed with error:", error.Message)
		t.FailNow()
	}
}

func AssertVError(t *testing.T, error *errors.VError, code errors.VErrorCode, message string) {
	if nil == error {
		logger.Error("Test failed because error is empty.")
		t.FailNow()
	}
	Assert(t, error.Code == code && error.Message == message, fmt.Sprintf("Error \n[%d - %s] \nwas expected but \n[%d - %s] \nwas found", code, message, error.Code, error.Message))
}

func NewTestDatabase(uuid string) (*sql.DB, *errors.VError) {

	path := fmt.Sprintf("%s/%s", TestDatabasePath, "temp")
	err := os.MkdirAll(path, 0775)
	if err != nil {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	name := fmt.Sprintf("%s/%s.db", path, uuid)
	db, err := database.Connect(name)
	if err != nil {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	err = tables.UpdateDatabaseTablesToLatestVersion(TestDatabasePath, tables.MainDatabase, db)
	if err != nil {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	return db, nil
}

func RemoveDatabase(uuid string) {
	path := fmt.Sprintf("%s/%s", TestDatabasePath, "temp")
	os.Remove(fmt.Sprintf("%s/%s.db", path, uuid))
}
