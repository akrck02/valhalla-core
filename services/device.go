package services

import (
	"database/sql"

	"github.com/akrck02/valhalla-core/database/dal"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func GetDevice(db *sql.DB, userID int64, userAgent string, address string) (*models.Device, *verrors.VError) {
	return dal.GetDevice(db, userID, userAgent, address)
}

func GetUserDevices(db *sql.DB, userID int64) ([]models.Device, *verrors.VError) {
	return dal.GetUserDevices(db, userID)
}

func DeleteDevice(db *sql.DB, userID int64, userAgent string, address string) *verrors.VError {
	return dal.DeleteDevice(db, userID, userAgent, address)
}
