package dal

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func RegisterUser(db *sql.DB, user *models.User, auth *models.UserAuth) (*int64, *errors.VError) {

	if nil == db {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	if nil == user {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "User cannot be empty.",
		}
	}

	if "" == strings.TrimSpace(user.Email) {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "User email cannot be empty.",
		}
	}

	if nil == auth {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "User auth cannot be empty.",
		}
	}

	if "" == strings.TrimSpace(auth.Password) {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "User password cannot be empty.",
		}
	}

	statement, err := db.Prepare("INSERT INTO user(email,profile_pic,insert_date) VALUES(?,?,?)")
	if nil != err {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.DatabaseError,
			Message: err.Error(),
		}
	}

	res, err := statement.Exec(user.Email, user.ProfilePicture, time.Now())
	if nil != err {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.DatabaseError,
			Message: err.Error(),
		}
	}

	user.Id, err = res.LastInsertId()
	if nil != err {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.DatabaseError,
			Message: err.Error(),
		}
	}

	return &user.Id, nil
}

func GetUser(db *sql.DB, id int64) (*models.User, *errors.VError) {

	if nil == db {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	if 0 < id {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "User id must be positive.",
		}
	}

	statement, err := db.Prepare(`
		SELECT email,
			profile_pic,
			insert_date
		FROM user
		WHERE id = ?
	`)

	if nil != err {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.DatabaseError,
			Message: err.Error(),
		}
	}

	rows, err := statement.Query(id)
	if nil != err {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.DatabaseError,
			Message: err.Error(),
		}
	}
	defer rows.Close()

	if rows.Next() {

		var email string
		var profilePic string
		var insertDate int64

		rows.Scan(&email, profilePic, insertDate)

		return &models.User{
			Id:             id,
			Email:          email,
			ProfilePicture: profilePic,
			InsertDate:     insertDate,
		}, nil
	}

	return nil, &errors.VError{
		Status:  http.StatusNotFound,
		Code:    errors.NotFound,
		Message: "User not found",
	}
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, *errors.VError) {

	if nil == db {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	if "" == strings.TrimSpace(email) {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "User id must be positive.",
		}
	}

	statement, err := db.Prepare(`
		SELECT id,
			profile_pic,
			insert_date
		FROM user
		WHERE email = ?
	`)

	if nil != err {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.DatabaseError,
			Message: err.Error(),
		}
	}

	rows, err := statement.Query(email)
	if nil != err {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.DatabaseError,
			Message: err.Error(),
		}
	}
	defer rows.Close()

	if rows.Next() {
		var id int64
		var profilePic string
		var insertDate int64

		rows.Scan(&email, profilePic, insertDate)

		return &models.User{
			Id:             id,
			Email:          email,
			ProfilePicture: profilePic,
			InsertDate:     insertDate,
		}, nil
	}

	return nil, &errors.VError{
		Status:  http.StatusNotFound,
		Code:    errors.NotFound,
		Message: "User not found",
	}
}

func DeleteUser(db *sql.DB, id string) *errors.VError {

	if nil == db {
		return &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	return errors.TODO()
}

func UpdateUser(db *sql.DB, id string, user string) *errors.VError {

	if nil == db {
		return &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	return errors.TODO()
}

func UpdateUserProfilePicture(db *sql.DB, id string, picture []byte) *errors.VError {

	if nil == db {
		return &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	return errors.TODO()
}

func Login(db *sql.DB, email string, device string) (*string, *errors.VError) {
	if nil == db {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	return nil, errors.TODO()
}

func LoginWithAuth(db *sql.DB, id string, token string) (*string, *errors.VError) {

	if nil == db {
		return nil, &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	return nil, errors.TODO()
}

func ValidateUserAccount(db *sql.DB, code string) *errors.VError {

	if nil == db {
		return &errors.VError{
			Status:  http.StatusInternalServerError,
			Code:    errors.InvalidRequest,
			Message: "Database connection cannot be empty.",
		}
	}

	return errors.TODO()
}
