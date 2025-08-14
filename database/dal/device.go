package dal

import (
	"database/sql"
	"time"

	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/logger"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func CreateDevice(db *sql.DB, userID int64, device *models.Device) *verrors.VError {
	statement, err := db.Prepare("INSERT INTO device(user_id, user_agent, address, token, insert_date, update_date) values(?,?,?,?,?,?)")
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(userID, device.UserAgent, device.Address, device.Token, time.Now(), time.Now())
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if 0 >= affectedRows {
		return verrors.New(verrors.NothingChangedErrorCode, "Device was not created.")
	}

	return nil
}

func GetDevice(db *sql.DB, userID int64, userAgent string, address string) (*models.Device, *verrors.VError) {
	statement, err := db.Prepare("SELECT user_agent, address, token, insert_date, update_date FROM device WHERE user_id = ? AND user_agent = ? AND address = ?")
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	rows, err := statement.Query(userID, userAgent, address)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, verrors.NotFound("Device not found.")
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

func GetUserDevices(db *sql.DB, userID int64) ([]models.Device, *verrors.VError) {

	statement, err := db.Prepare("SELECT user_agent, address, token, insert_date, update_date FROM device WHERE user_id = ?")
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	rows, err := statement.Query(userID)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	devices := []models.Device{}
	for rows.Next() {

		var device models.Device
		rows.Scan(&device)
		logger.Log(device)
		devices = append(devices, device)
	}

	return devices, nil
}

func GetDeviceByAuth(db *sql.DB, userID int64, userAgent string, address string, token string) (*models.Device, *verrors.VError) {
	statement, err := db.Prepare("SELECT user_agent, address, token, insert_date, update_date FROM device WHERE user_id = ? AND user_agent = ? AND address = ? AND token = ?")
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	rows, err := statement.Query(userID, userAgent, address, token)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, verrors.NotFound("Device not found.")
	}

	var d models.Device
	rows.Scan(&d)

	return &d, nil
}

func UpdateDeviceToken(db *sql.DB, userID int64, userAgent string, address string, token string) *verrors.VError {
	statement, err := db.Prepare("UPDATE device SET token = ?, update_date = ? WHERE user_id = ? AND user_agent = ? AND address = ?")
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(token, time.Now(), userID, userAgent, address)
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if 0 >= affectedRows {
		return verrors.NotFound("Device does not exist.")
	}

	return nil
}

func DeleteDevice(db *sql.DB, userID int64, userAgent string, address string) *verrors.VError {
	statement, err := db.Prepare("DELETE FROM device WHERE user_id = ? AND user_agent = ? AND address = ?")
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(userID, userAgent, address)
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if 0 >= affectedRows {
		return verrors.NotFound("Device does not exist.")
	}

	return nil
}
