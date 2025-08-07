package dal

import (
	"database/sql"
	"strings"
	"time"

	"github.com/akrck02/valhalla-core/sdk/cryptography"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/akrck02/valhalla-core/sdk/validations"
	"github.com/google/uuid"
)

func RegisterUser(db *sql.DB, user *models.User) (*int64, *errors.VError) {

	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	if nil == user {
		return nil, errors.New(errors.InvalidRequest, "User cannot be empty.")
	}

	err := validations.ValidateEmail(user.Email)
	if nil != err {
		return nil, errors.New(errors.InvalidEmail, err.Error())
	}

	err = validations.ValidatePassword(user.Password)
	if nil != err {
		return nil, errors.New(errors.InvalidPassword, err.Error())
	}

	hashedPassword, err := cryptography.Hash(user.Password)
	if nil != err {
		return nil, errors.Unexpected(err.Error())
	}

	statement, err := db.Prepare("INSERT INTO user(email, profile_pic, password, database, validation_code, insert_date) VALUES(?,?,?,?,?,?)")
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(user.Email, user.ProfilePicture, hashedPassword, uuid.NewString(), uuid.NewString(), time.Now())
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	user.Id, err = res.LastInsertId()
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	return &user.Id, nil
}

func GetUser(db *sql.DB, id int64) (*models.User, *errors.VError) {

	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	if 0 >= id {
		return nil, errors.New(errors.InvalidId, "User id must be positive.")
	}

	statement, err := db.Prepare(`
		SELECT id,
			email,
			profile_pic,
			password,
			database,
			insert_date
		FROM user
		WHERE id = ?
	`)

	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(id)
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}
	defer rows.Close()

	if rows.Next() {

		var id int64
		var email string
		var profilePicture string
		var password string
		var database string
		var insertDate int64

		rows.Scan(
			&id,
			&email,
			&profilePicture,
			&password,
			&database,
			&insertDate,
		)

		return &models.User{
			Id:             id,
			Email:          email,
			ProfilePicture: profilePicture,
			Password:       password,
			Database:       database,
			InsertDate:     insertDate,
		}, nil

	}

	return nil, errors.New(errors.NotFound, "User not found.")
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, *errors.VError) {

	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	if "" == strings.TrimSpace(email) {
		return nil, errors.New(errors.InvalidId, "User id must be positive")
	}

	statement, err := db.Prepare(`
		SELECT id,
			email,
			profile_pic,
			password,
			database,
			insert_date
		FROM user
		WHERE email = ?
	`)

	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(email)
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}
	defer rows.Close()

	if rows.Next() {

		var id int64
		var email string
		var profilePicture string
		var password string
		var database string
		var insertDate int64

		rows.Scan(
			&id,
			&email,
			&profilePicture,
			&password,
			&database,
			&insertDate,
		)

		return &models.User{
			Id:             id,
			Email:          email,
			ProfilePicture: profilePicture,
			Password:       password,
			Database:       database,
			InsertDate:     insertDate,
		}, nil

	}

	return nil, errors.New(errors.NotFound, "User not found.")
}

func DeleteUser(db *sql.DB, id int64) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	if 0 >= id {
		return errors.New(errors.InvalidId, "User id must be positive.")
	}

	statement, err := db.Prepare("DELETE FROM user WHERE id=?")
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(id)
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	if 0 == affectedRows {
		return errors.New(errors.DatabaseError, "Cannot delete user.")
	}

	return nil
}

func UpdateUserEmail(db *sql.DB, id int64, email string) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	if 0 >= id {
		return errors.New(errors.InvalidId, "User id must be positive.")
	}

	err := validations.ValidateEmail(email)
	if nil != err {
		return errors.New(errors.InvalidEmail, err.Error())
	}

	statement, err := db.Prepare("UPDATE user SET email = ? WHERE id = ?")
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(email, id)
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	if 0 == affectedRows {
		return errors.New(errors.DatabaseError, "Cannot update user.")
	}

	return nil
}

func UpdateUserProfilePicture(db *sql.DB, id int64, profilePic string) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	if 0 >= id {
		return errors.New(errors.InvalidId, "User id must be positive.")
	}

	statement, err := db.Prepare("UPDATE user SET profile_pic = ? WHERE id = ?")

	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(profilePic, id)
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return errors.New(errors.DatabaseError, err.Error())
	}

	if 0 == affectedRows {
		return errors.New(errors.DatabaseError, "Cannot update user.")
	}

	return nil
}

func Login(db *sql.DB, email string, device *models.Device) (*string, *errors.VError) {

	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	if "" == email {
		return nil, errors.New(errors.InvalidEmail, "Email cannot be empty.")
	}

	if nil == device {
		return nil, errors.New(errors.InvalidRequest, "Device cannot be empty.")
	}

	if "" == device.Address {
		return nil, errors.New(errors.InvalidRequest, "Device address cannot be empty.")
	}

	if "" == device.UserAgent {
		return nil, errors.New(errors.InvalidRequest, "Device user agent cannot be empty.")
	}

	return nil, nil
}

func LoginWithAuth(db *sql.DB, id int64, token string) (*string, *errors.VError) {

	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	return nil, nil
}

func ValidateUserAccount(db *sql.DB, code string) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	return nil
}
