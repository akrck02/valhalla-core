package dal

import (
	"database/sql"
	"time"

	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func CreateDevice(db *sql.DB, userId int64, device *models.Device) *errors.VError {

	statement, err := db.Prepare("INSERT INTO device(user_id, user_agent, address, token, insert_date, update_date) values(?,?,?,?,?,?)")
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(userId, device.UserAgent, device.Address, device.Token, time.Now(), time.Now())
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	if 0 >= affectedRows {
		return errors.New(errors.NothingChanged, "Device was not created.")
	}

	return nil
}

func UpdateDeviceToken(db *sql.DB, userId int64, userAgent string, address string, token string) *errors.VError {

	return nil
}

func DeleteDevice(db *sql.DB, userId int64, userAgent string, address string) *errors.VError {

	return nil
}
