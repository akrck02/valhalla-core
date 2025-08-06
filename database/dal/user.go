package dal

import (
	"database/sql"
	"strings"
	"time"

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

	statement, err := db.Prepare("INSERT INTO user(email, profile_pic, password, database, insert_date) VALUES(?,?,?,?,?)")
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(user.Email, user.ProfilePicture, user.Password, uuid.NewString(), time.Now())
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

	if 0 > id {
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
		user := &models.User{}
		rows.Scan(
			user.Id,
			user.Email,
			user.ProfilePicture,
			user.Password,
			user.Database,
			user.InsertDate,
		)
		return user, nil
	}

	return nil, errors.New(errors.NotFound, "User not found")
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
		user := &models.User{}
		rows.Scan(
			user.Id,
			user.Email,
			user.ProfilePicture,
			user.Password,
			user.Database,
			user.InsertDate,
		)
		return user, nil
	}

	return nil, errors.New(errors.NotFound, "User not found")
}

func DeleteUser(db *sql.DB, id string) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	return nil
}

func UpdateUser(db *sql.DB, id string, user string) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	return nil
}

func UpdateUserProfilePicture(db *sql.DB, id string, picture []byte) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	return nil
}

func Login(db *sql.DB, email string, device string) (*string, *errors.VError) {
	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	return nil, nil
}

func LoginWithAuth(db *sql.DB, id string, token string) (*string, *errors.VError) {

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
