package tests

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/tables"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/logger"
)

const TestDatabasePath string = "../../.."

func Assert(t *testing.T, predicate bool, failMessage string) {
	if !predicate {
		logger.Log("Test failed:", failMessage)
		t.FailNow()
	}
}

func AssertVErrorDoesNotExist(t *testing.T, error *verrors.VError) {
	if nil != error {
		logger.Error("Test failed with error:", error.Message)
		t.FailNow()
	}
}

func AssertVError(t *testing.T, error *verrors.VError, code verrors.VErrorCode, message string) {
	if nil == error {
		logger.Error("Test failed because error is empty.")
		t.FailNow()
	}
	Assert(t, error.Code == code && error.Message == message, fmt.Sprintf("\n[%d - %s] \nwas expected but \n[%d - %s] \nwas found\n", code, message, error.Code, error.Message))
}

func NewTestDatabase() (*sql.DB, *verrors.VError) {

	db, err := database.Connect(":memory:")
	if err != nil {
		return nil, verrors.New(verrors.DatabaseErrorCode, err.Error())
	}

	err = tables.UpdateDatabaseTablesToLatestVersion(TestDatabasePath, tables.MainDatabase, db)
	if err != nil {
		return nil, verrors.New(verrors.DatabaseErrorCode, err.Error())
	}

	return db, nil
}
