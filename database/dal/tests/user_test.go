package tests

import (
	"database/sql"
	"testing"

	"github.com/akrck02/valhalla-core/database/dal"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/google/uuid"
)

func TestUserCrud(t *testing.T) {

	databaseUuid := uuid.NewString()
	db, err := NewTestDatabase(databaseUuid)
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()
	defer RemoveDatabase(databaseUuid)

	expectedUser := &models.User{
		Email:    "user@valhalla.org",
		Password: "$P4ssw0rdW3db1128",
	}

	t.Run("Register validations", func(t *testing.T) { registerValidations(t, db) })
	t.Run("Register user", func(t *testing.T) { expectedUser = registerUser(t, db, expectedUser) })

	t.Run("Get user validations", func(t *testing.T) { getUserValidations(t, db) })

	t.Run("Update user mail validation errors", func(t *testing.T) { updateEmailValidations(t, db) })
	t.Run("Update user mail", func(t *testing.T) { expectedUser = updateUserEmail(t, db, expectedUser) })

	t.Run("Update user profile picture validation errors", func(t *testing.T) { updateUserProfilePictureValidations(t, db) })
	t.Run("Update user profile", func(t *testing.T) { expectedUser = updateUserProfilePicture(t, db, expectedUser) })

}

func registerValidations(t *testing.T, db *sql.DB) {

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

	obtainedUser, err := dal.GetUser(db, *userId)
	AssertVErrorDoesNotExist(t, err)
	Assert(t, nil != obtainedUser && obtainedUser.Email == obtainedUser.Email, "Expected user and obtained user mismatch")
	return obtainedUser
}

func getUserValidations(t *testing.T, db *sql.DB) {
	_, err := dal.GetUser(db, 0)
	AssertVError(t, err, errors.InvalidId, "User id must be positive.")

	_, err = dal.GetUser(db, 999)
	AssertVError(t, err, errors.NotFound, "User not found.")
}

func updateEmailValidations(t *testing.T, db *sql.DB) {
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

func updateUserEmail(t *testing.T, db *sql.DB, user *models.User) *models.User {

	newMail := "user-modified@valhalla.org"
	userId := user.Id
	err := dal.UpdateUserEmail(db, userId, newMail)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, userId)
	AssertVErrorDoesNotExist(t, err)

	Assert(t, obtainedUser.Email == newMail, "User mail mismatch.")
	return obtainedUser
}

func updateUserProfilePictureValidations(t *testing.T, db *sql.DB) {
	err := dal.UpdateUserProfilePicture(nil, 0, "")
	AssertVError(t, err, errors.UnexpectedError, "Database connection cannot be empty.")

	err = dal.UpdateUserProfilePicture(db, 0, "")
	AssertVError(t, err, errors.InvalidId, "User id must be positive.")
}

func updateUserProfilePicture(t *testing.T, db *sql.DB, user *models.User) *models.User {

	newProfilePic := "my-user-profile-pic.jpg"
	userId := user.Id
	err := dal.UpdateUserProfilePicture(db, userId, newProfilePic)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, userId)
	AssertVErrorDoesNotExist(t, err)

	Assert(t, obtainedUser.ProfilePicture == newProfilePic, "User mail mismatch.")
	return obtainedUser
}

func TestUserLogin(t *testing.T) {

	databaseUuid := uuid.NewString()
	db, err := NewTestDatabase(databaseUuid)
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()
	defer RemoveDatabase(databaseUuid)

	expectedUser := &models.User{
		Email:    "user@valhalla.org",
		Password: "$P4ssw0rdW3db1128",
	}

	userId, err := dal.RegisterUser(db, expectedUser)
	expectedUser.Id = *userId

	t.Run("Login validation errors", func(t *testing.T) { loginValidation(t, db) })
	t.Run("Login", func(t *testing.T) { login(t, db) })

	t.Run("Login with auth validation errors", func(t *testing.T) { loginWithAuthValidation(t, db) })
	t.Run("Login with auth", func(t *testing.T) { loginWithAuth(t, db) })

	t.Run("Validate user account validation errors", func(t *testing.T) { validateUserAccountValidation(t, db) })
	t.Run("Validate user account", func(t *testing.T) { validateUserAccount(t, db) })

}

func loginValidation(t *testing.T, db *sql.DB) {

}

func login(t *testing.T, db *sql.DB) {

}

func loginWithAuthValidation(t *testing.T, db *sql.DB) {

}

func loginWithAuth(t *testing.T, db *sql.DB) {

}

func validateUserAccountValidation(t *testing.T, db *sql.DB) {

}

func validateUserAccount(t *testing.T, db *sql.DB) {

}
