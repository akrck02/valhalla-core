package tests

import (
	"testing"

	"github.com/akrck02/valhalla-core/database/dal"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/logger"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func TestRegister(t *testing.T) {

	db, err := NewTestDatabase()
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()

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

	expectedUser := &models.User{
		Email:    "akrck02@gmail.com",
		Password: "$P4ssw0rdW3db1128",
	}
	userId, err := dal.RegisterUser(db, expectedUser)
	AssertVErrorDoesNotExist(t, err)

	user, err := dal.GetUser(db, *userId)
	logger.Log(expectedUser, user)
	AssertVErrorDoesNotExist(t, err)
	Assert(t, nil != user && user.Email == expectedUser.Email, "Expected user and obtained user mismatch")

}
