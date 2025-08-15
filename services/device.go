package services

import (
	"database/sql"

	"github.com/akrck02/valhalla-core/database/dal"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func GetDevice(db *sql.DB, userID int64, userAgent string, address string) (*models.Device, *verrors.VError) {
	device, err := dal.GetDevice(db, userID, userAgent, address)
	if nil != err {
		return nil, err
	}

	return (*models.Device)(device), nil
}

func GetUserDevices(db *sql.DB, userID int64) ([]*models.Device, *verrors.VError) {
	obtainedDevices, err := dal.GetUserDevices(db, userID)
	if nil != err {
		return nil, err
	}

	var devices = []*models.Device{}
	for _, device := range obtainedDevices {
		devices = append(devices, (*models.Device)(device))
	}

	return devices, nil
}

func DeleteDevice(db *sql.DB, userID int64, userAgent string, address string) *verrors.VError {
	return dal.DeleteDevice(db, userID, userAgent, address)
}
