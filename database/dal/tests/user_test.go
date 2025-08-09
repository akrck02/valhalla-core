package tests

import (
	"database/sql"
	"testing"

	"github.com/akrck02/valhalla-core/database/dal"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/google/uuid"
)

func TestUserCrud(t *testing.T) {
	databaseUUID := uuid.NewString()
	db, err := NewTestDatabase(databaseUUID)
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()
	defer RemoveDatabase(databaseUUID)

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
	AssertVError(t, err, verrors.UnexpectedError, verrors.DatabaseConnectionEmptyMessage)

	_, err = dal.RegisterUser(db, nil)
	AssertVError(t, err, verrors.InvalidRequest, verrors.UserEmptyMessage)

	_, err = dal.RegisterUser(db, &models.User{})
	AssertVError(t, err, verrors.InvalidEmail, verrors.EmailEmptyMessage)

	_, err = dal.RegisterUser(db, &models.User{Email: "uservalhalla.org"})
	AssertVError(t, err, verrors.InvalidEmail, "mail: missing '@' or angle-addr")

	_, err = dal.RegisterUser(db, &models.User{Email: "u@.org"})
	AssertVError(t, err, verrors.InvalidEmail, "mail: missing '@' or angle-addr")

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org"})
	AssertVError(t, err, verrors.InvalidPassword, verrors.PasswordEmptyMessage)

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: ""})
	AssertVError(t, err, verrors.InvalidPassword, verrors.PasswordEmptyMessage)

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "abcdefghijklmnñopq"})
	AssertVError(t, err, verrors.InvalidPassword, verrors.PasswordNoNumericMessage)

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "1bcdefghijklmnñopq"})
	AssertVError(t, err, verrors.InvalidPassword, verrors.PasswordNoSpecialCharacterMessage)

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "#1bcdefghijklmnñop"})
	AssertVError(t, err, verrors.InvalidPassword, verrors.PasswordNoUppercaseCharacterMessage)

	_, err = dal.RegisterUser(db, &models.User{Email: "user@valhalla.org", Password: "#1BCDEFGHUJKLMNÑOP"})
	AssertVError(t, err, verrors.InvalidPassword, verrors.PasswordNoLowercaseCharacterMessage)
}

func registerUser(t *testing.T, db *sql.DB, user *models.User) *models.User {
	userID, err := dal.RegisterUser(db, user)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, *userID)
	AssertVErrorDoesNotExist(t, err)
	Assert(t, nil != obtainedUser && user.Email == obtainedUser.Email, "expected user and obtained user mismatch")
	return obtainedUser
}

func getUserValidations(t *testing.T, db *sql.DB) {
	_, err := dal.GetUser(db, 0)
	AssertVError(t, err, verrors.InvalidID, verrors.UserIDNegativeMessage)

	_, err = dal.GetUser(db, 999)
	AssertVError(t, err, verrors.NotFound, verrors.UserNotFoundMessage)
}

func updateEmailValidations(t *testing.T, db *sql.DB) {
	err := dal.UpdateUserEmail(nil, 0, "")
	AssertVError(t, err, verrors.UnexpectedError, verrors.DatabaseConnectionEmptyMessage)

	err = dal.UpdateUserEmail(db, 0, "")
	AssertVError(t, err, verrors.InvalidID, verrors.UserIDNegativeMessage)

	err = dal.UpdateUserEmail(db, 1, "")
	AssertVError(t, err, verrors.InvalidEmail, verrors.EmailEmptyMessage)

	err = dal.UpdateUserEmail(db, 1, "uservalhalla.org")
	AssertVError(t, err, verrors.InvalidEmail, "mail: missing '@' or angle-addr")

	err = dal.UpdateUserEmail(db, 1, "u@.org")
	AssertVError(t, err, verrors.InvalidEmail, "mail: missing '@' or angle-addr")
}

func updateUserEmail(t *testing.T, db *sql.DB, user *models.User) *models.User {
	newMail := "user-modified@valhalla.org"
	userID := user.ID
	err := dal.UpdateUserEmail(db, userID, newMail)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, userID)
	AssertVErrorDoesNotExist(t, err)

	Assert(t, obtainedUser.Email == newMail, "user mail mismatch.")
	return obtainedUser
}

func updateUserProfilePictureValidations(t *testing.T, db *sql.DB) {
	err := dal.UpdateUserProfilePicture(nil, 0, "")
	AssertVError(t, err, verrors.UnexpectedError, verrors.DatabaseConnectionEmptyMessage)

	err = dal.UpdateUserProfilePicture(db, 0, "")
	AssertVError(t, err, verrors.InvalidID, verrors.UserIDNegativeMessage)
}

func updateUserProfilePicture(t *testing.T, db *sql.DB, user *models.User) *models.User {
	newProfilePic := "my-user-profile-pic.jpg"
	userID := user.ID
	err := dal.UpdateUserProfilePicture(db, userID, newProfilePic)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, userID)
	AssertVErrorDoesNotExist(t, err)

	Assert(t, obtainedUser.ProfilePicture == newProfilePic, "user mail mismatch.")
	return obtainedUser
}

func TestUserLogin(t *testing.T) {
	databaseUUID := uuid.NewString()
	db, err := NewTestDatabase(databaseUUID)
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()
	defer RemoveDatabase(databaseUUID)

	expectedUser := &models.User{
		Email:    "user@valhalla.org",
		Password: "$P4ssw0rdW3db1128",
	}

	userID, err := dal.RegisterUser(db, expectedUser)
	AssertVErrorDoesNotExist(t, err)
	expectedUser.ID = *userID

	t.Run("Login validation errors", func(t *testing.T) { loginValidation(t, db) })
	t.Run("Login", func(t *testing.T) { login(t, db) })

	t.Run("Login with auth validation errors", func(t *testing.T) { loginWithAuthValidation(t, db) })
	t.Run("Login with auth", func(t *testing.T) { loginWithAuth(t, db) })

	t.Run("Validate user account validation errors", func(t *testing.T) { validateUserAccountValidation(t, db) })
	t.Run("Validate user account", func(t *testing.T) { validateUserAccount(t, db) })
}

func loginValidation(t *testing.T, db *sql.DB) {
	_, err := dal.Login(nil, "", []string{}, "", "", "", nil)
	AssertVError(t, err, verrors.UnexpectedError, verrors.DatabaseConnectionEmptyMessage)

	_, err = dal.Login(db, "", []string{}, "", "", "", nil)
	AssertVError(t, err, verrors.UnexpectedError, verrors.ServiceIDEmptyMessage)

	_, err = dal.Login(db, "valhalla-core", []string{}, "", "", "", nil)
	AssertVError(t, err, verrors.UnexpectedError, verrors.RegisteredDomainsEmptyMessage)

	_, err = dal.Login(db, "valhalla-core", []string{"https://valhalla.org"}, "", "", "", nil)
	AssertVError(t, err, verrors.UnexpectedError, verrors.SecretEmptyMessage)

	_, err = dal.Login(db, "valhalla-core", []string{"https://valhalla.org"}, "secret", "", "", nil)
	AssertVError(t, err, verrors.InvalidEmail, verrors.EmailEmptyMessage)

	_, err = dal.Login(db, "valhalla-core", []string{"https://valhalla.org"}, "secret", "user@valhalla.org", "", nil)
	AssertVError(t, err, verrors.InvalidRequest, verrors.DeviceEmptyMessage)

	_, err = dal.Login(db, "valhalla-core", []string{"https://valhalla.org"}, "secret", "user@valhalla.org", "", &models.Device{})
	AssertVError(t, err, verrors.InvalidRequest, verrors.DeviceAddressEmptyMessage)

	_, err = dal.Login(db, "valhalla-core", []string{"https://valhalla.org"}, "secret", "user@valhalla.org", "", &models.Device{Address: "127.0.0.1"})
	AssertVError(t, err, verrors.InvalidRequest, verrors.DeviceUserAgentEmptyMessage)
}

func login(t *testing.T, db *sql.DB) {
	token, err := dal.Login(
		db,
		"valhalla-core",
		[]string{"https://valhalla.org"},
		"secret",
		"user@valhalla.org",
		"$P4ssw0rdW3db1128",
		&models.Device{
			Address:   "127.0.0.1",
			UserAgent: "Firefox",
		},
	)
	AssertVErrorDoesNotExist(t, err)

	device, err := dal.GetDevice(db, 1, "Firefox", "127.0.0.1")
	AssertVErrorDoesNotExist(t, err)
	Assert(t, device.Token == *token, "token mismatch")

	token, err = dal.Login(
		db,
		"valhalla-core",
		[]string{"https://valhalla.org"},
		"secret",
		"user@valhalla.org",
		"$P4ssw0rdW3db1128",
		&models.Device{
			Address:   "127.0.0.1",
			UserAgent: "Firefox",
		},
	)
	AssertVErrorDoesNotExist(t, err)

	device, err = dal.GetDevice(db, 1, "Firefox", "127.0.0.1")
	AssertVErrorDoesNotExist(t, err)
	Assert(t, device.Token == *token, "token mismatch")
}

func loginWithAuthValidation(t *testing.T, db *sql.DB) {
	err := dal.LoginWithAuth(nil, "", "")
	AssertVError(t, err, verrors.UnexpectedError, verrors.DatabaseConnectionEmptyMessage)

	err = dal.LoginWithAuth(db, "", "")
	AssertVError(t, err, verrors.UnexpectedError, verrors.SecretEmptyMessage)

	err = dal.LoginWithAuth(db, "secret", "")
	AssertVError(t, err, verrors.InvalidToken, verrors.TokenEmptyMessage)

	err = dal.LoginWithAuth(db, "secret", "token")
	AssertVError(t, err, verrors.InvalidToken, "token is malformed: token contains an invalid number of segments")
}

func loginWithAuth(t *testing.T, db *sql.DB) {
	token, err := dal.Login(
		db,
		"valhalla-core",
		[]string{"https://valhalla.org"},
		"secret",
		"user@valhalla.org",
		"$P4ssw0rdW3db1128",
		&models.Device{
			Address:   "127.0.0.1",
			UserAgent: "Firefox",
		},
	)
	AssertVErrorDoesNotExist(t, err)

	err = dal.LoginWithAuth(db, "secret", *token)
	AssertVErrorDoesNotExist(t, err)

	err = dal.LoginWithAuth(db, "secret1", *token)
	AssertVError(t, err, verrors.InvalidToken, "token signature is invalid: signature is invalid")
}

func validateUserAccountValidation(t *testing.T, db *sql.DB) {
}

func validateUserAccount(t *testing.T, db *sql.DB) {
}
