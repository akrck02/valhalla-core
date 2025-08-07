package tests

import (
	"testing"

	"github.com/akrck02/valhalla-core/database/dal"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/google/uuid"
)

func TestRegister(t *testing.T) {

	databaseUuid := uuid.NewString()
	db, err := NewTestDatabase(databaseUuid)
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()
	defer RemoveDatabase(databaseUuid)

	t.Run("Validation errors", func(t *testing.T) {
		_, err = dal.RegisterUser(nil, nil)
		AssertVError(t, err, errors.UnexpectedError, "Database connection cannot be empty.")

		_, err = dal.RegisterUser(db, nil)
		AssertVError(t, err, errors.InvalidRequest, "User cannot be empty.")

		_, err = dal.RegisterUser(db, &models.User{})
		AssertVError(t, err, errors.InvalidEmail, "Email cannot be empty.")

		_, err = dal.RegisterUser(db, &models.User{Email: "uservalhalla.org"})
		AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")

		_, err = dal.RegisterUser(db, &models.User{Email: "u@.org"})
		AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")
		ot be empty.]
was found
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

	})

	t.Run("Register user logic", func(t *testing.T) {

		expectedUser := &models.User{
			Email:    "user@valhalla.org",
			Password: "$P4ssw0rdW3db1128",
		}
		userId, err := dal.RegisterUser(db, expectedUser)
		AssertVErrorDoesNotExist(t, err)

		t.Run("Get user validations", func(t *testing.T) {
			_, err := dal.GetUser(db, 0)
			AssertVError(t, err, errors.InvalidId, "User id must be positive.")

			_, err = dal.GetUser(db, 999)
			AssertVError(t, err, errors.NotFound, "User not found.")
		})

		user, err := dal.GetUser(db, *userId)
		AssertVErrorDoesNotExist(t, err)
		Assert(t, nil != user && user.Email == expectedUser.Email, "Expected user and obtained user mismatch")

	})

}

func TestUpdateUserEmail(t *testing.T) {

	databaseUuid := uuid.NewString()
	db, err := NewTestDatabase(databaseUuid)
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()
	defer RemoveDatabase(databaseUuid)

	t.Run("Validation errors", func(t *testing.T) {

		err = dal.UpdateUserEmail(nil, 0, "")
		AssertVError(t, err, errors.UnexpectedError, "Database connection cannot be empty.")

		err = dal.UpdateUserEmail(db, 0, "")
		AssertVError(t, err, errors.InvalidId, "User id must be positive.")

		err = dal.UpdateUserEmail(db, 1, "")
		AssertVError(t, err, errors.InvalidEmail, "Email cannot be empty.")

		err = dal.UpdateUserEmail(db, 1, "uservalhalla.org")
		AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")

		err = dal.UpdateUserEmail(db, 1, "u@.org")
		AssertVError(t, err, errors.InvalidEmail, "mail: missing '@' or angle-addr")
	})

	t.Run("Update user mail logic", func(t *testing.T) {

		expectedUser := &models.User{
			Email:    "user@valhalla.org",
			Password: "$P4ssw0rdW3db1128",
		}
		userId, err := dal.RegisterUser(db, expectedUser)
		AssertVErrorDoesNotExist(t, err)

		newMail := "user-modified@valhalla.org"
		err = dal.UpdateUserEmail(db, *userId, newMail)
		AssertVErrorDoesNotExist(t, err)

		user, err := dal.GetUser(db, *userId)
		AssertVErrorDoesNotExist(t, err)

		Assert(t, user.Email == newMail, "User mail mismatch.")
	})
}
