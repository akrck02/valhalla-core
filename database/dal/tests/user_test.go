package tests

import (
	"database/sql"
	"testing"

	"github.com/akrck02/valhalla-core/database/dal"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/google/uuid"
)

func UserTest(t *testing.T) {

	databaseUuid := uuid.NewString()
	db, err := NewTestDatabase(databaseUuid)
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()
	defer RemoveDatabase(databaseUuid)

	expectedUser := &models.User{
		Email:    "user@valhalla.org",
		Password: "$P4ssw0rdW3db1128",
	}

	t.Run("Register validations", func(t *testing.T) { registerValidationErrors(t, db) })
	t.Run("Register user logic", func(t *testing.T) { expectedUser = registerUser(t, db, expectedUser) })
	t.Run("Update user mail validation errors", func(t *testing.T) { updateEmailValidateError(t, db) })
	t.Run("Update user mail logic", func(t *testing.T) { expectedUser = updateUserEmail(t, db, expectedUser) })

}

func registerValidationErrors(t *testing.T, db *sql.DB) {

	_, err := dal.RegisterUser(nil, nil)
	AssertVError(t, err, errors.UnexpectedError, "Database connection cannot be empty.")

	_, err = dal.RegisterUser(db, nil)
	AssertVError(t, err, errors.InvalidRequest, "User cannot be empty.")

	_, err = dal.RegisterUser(db, &models.User{})
	AssertVError(t, err, errors.InvalidEmail, "Email cannot be empty.")

	_, err = dal.RegisterUser(db, &models.User{Email: "uservalhalla.org"})
	AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")

	_, err = dal.RegisterUser(db, &models.User{Email: "u@.org"})
	AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org"})
	AssertVError(t, err, errors.InvalidPassword, "Password cannot be empty.")

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: ""})
	AssertVError(t, err, errors.InvalidPassword, "Password cannot be empty.")

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "abcdefghijklmnñopq"})
	AssertVError(t, err, errors.InvalidPassword, "Password must contain at least one numeric character.")

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "1bcdefghijklmnñopq"})
	AssertVError(t, err, errors.InvalidPassword, "Password must contain at least one special character.")

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "#1bcdefghijklmnñop"})
	AssertVError(t, err, errors.InvalidPassword, "Password must contain at least one uppercase character.")

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "#1BCDEFGHUJKLMNÑOP"})
	AssertVError(t, err, errors.InvalidPassword, "Password must contain at least one lowercase character.")
}

func registerUser(t *testing.T, db *sql.DB, user *models.User) *models.User {

	userId, err := dal.RegisterUser(db, user)
	AssertVErrorDoesNotExist(t, err)

	t.Run("Get user validations", func(t *testing.T) {
		_, err := dal.GetUser(db, 0)
		AssertVError(t, err, errors.InvalidId, "User id must be positive.")

		_, err = dal.GetUser(db, 999)
		AssertVError(t, err, errors.NotFound, "User not found.")
	})

	obtainedUser, err := dal.GetUser(db, *userId)
	AssertVErrorDoesNotExist(t, err)
	Assert(t, nil != obtainedUser && obtainedUser.Email == obtainedUser.Email, "Expected user and obtained user mismatch")
	return obtainedUser
}

func updateUserEmail(t *testing.T, db *sql.DB, user *models.User) *models.User {

	newMail := "user-modified@valhalla.org"
	userId := *&user.Id
	err := dal.UpdateUserEmail(db, userId, newMail)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, userId)
	AssertVErrorDoesNotExist(t, err)

	Assert(t, user.Email == newMail, "User mail mismatch.")
	return obtainedUser
}

func updateEmailValidateError(t *testing.T, db *sql.DB) {
	err := dal.UpdateUserEmail(nil, 0, "")
	AssertVError(t, err, errors.UnexpectedError, "Database connection cannot be empty.")

	err = dal.UpdateUserEmail(db, 0, "")
	AssertVError(t, err, errors.InvalidId, "User id must be positive.")

	err = dal.UpdateUserEmail(db, 1, "")
	AssertVError(t, err, errors.InvalidEmail, "Email cannot be empty.")

	err = dal.UpdateUserEmail(db, 1, "uservalhalla.org")
	AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")

	err = dal.UpdateUserEmail(db, 1, "u@.org")
	AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")
}
