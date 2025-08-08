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

func GetDevice(db *sql.DB, userId int64, userAgent string, address string) (*models.Device, *errors.VError) {

	statement, err := db.Prepare("SELECT user_agent, address, token, insert_date, update_date FROM device WHERE user_id = ? AND user_agent = ? AND address = ?")
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(userId, userAgent, address)
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New(errors.NotFound, "Device not found.")
	}

	var obtainedUserAgent string
	var obtainedAddress string
	var token string
	var insertDate int64
	var updateDate int64
	rows.Scan(&obtainedUserAgent, &obtainedAddress, &token, &insertDate, &updateDate)

	return &models.Device{
		UserAgent:  obtainedUserAgent,
		Address:    obtainedAddress,
		Token:      token,
		InsertDate: insertDate,
		UpdateDate: updateDate,
	}, nil
}

func UpdateDeviceToken(db *sql.DB, userId int64, userAgent string, address string, token string) *errors.VError {

	statement, err := db.Prepare("UPDATE device SET token = ? WHERE user_id = ? AND user_agent = ? AND address = ?")
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(token, userId, userAgent, address)
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	if 0 >= affectedRows {
		return errors.New(errors.NotFound, "Device does not exist.")
	}

	return nil
}

func DeleteDevice(db *sql.DB, userId int64, userAgent string, address string) *errors.VError {

	statement, err := db.Prepare("DELETE FROM device WHERE user_id = ? AND user_agent = ? AND address = ?")
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(userId, userAgent, address)
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	if 0 >= affectedRows {
		return errors.New(errors.NotFound, "Device does not exist.")
	}

	return nil
}
