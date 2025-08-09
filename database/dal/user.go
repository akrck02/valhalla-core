package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/akrck02/valhalla-core/sdk/cryptography"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/akrck02/valhalla-core/sdk/validations"
)

func RegisterUser(db *sql.DB, user *models.User) (*int64, *verrors.VError) {
	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if nil == user {
		return nil, verrors.New(verrors.InvalidRequest, verrors.UserEmptyMessage)
	}

	err := validations.ValidateEmail(user.Email)
	if nil != err {
		return nil, verrors.New(verrors.InvalidEmail, err.Error())
	}

	usr, userGetErr := GetUserByEmail(db, user.Email)
	if nil == userGetErr && nil != usr {
		return nil, verrors.New(verrors.UserAlreadyExists, verrors.UserAlreadyExistsMessage)
	}

	if userGetErr.Code != verrors.NotFound {
		return nil, userGetErr
	}

	err = validations.ValidatePassword(user.Password)
	if nil != err {
		return nil, verrors.New(verrors.InvalidPassword, err.Error())
	}

	hashedPassword, err := cryptography.Hash(user.Password)
	if nil != err {
		return nil, verrors.Unexpected(err.Error())
	}

	statement, err := db.Prepare(
		"INSERT INTO user(email, profile_pic, password, database, validation_code, insert_date) VALUES(?,?,?,?,?,?)",
	)

	if nil != err {
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(
		user.Email,
		user.ProfilePicture,
		hashedPassword,
		uuid.NewString(),
		uuid.NewString(),
		time.Now(),
	)

	if nil != err {
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}

	user.ID, err = res.LastInsertId()
	if nil != err {
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}

	return &user.ID, nil
}

func GetUser(db *sql.DB, id int64) (*models.User, *verrors.VError) {
	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if 0 >= id {
		return nil, verrors.New(verrors.InvalidID, verrors.UserIDNegativeMessage)
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
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(id)
	if nil != err {
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, verrors.New(verrors.NotFound, verrors.UserNotFoundMessage)
	}

	var obtainedID int64
	var email string
	var profilePicture string
	var password string
	var database string
	var insertDate int64

	rows.Scan(
		&obtainedID,
		&email,
		&profilePicture,
		&password,
		&database,
		&insertDate,
	)

	return &models.User{
		ID:             obtainedID,
		Email:          email,
		ProfilePicture: profilePicture,
		Password:       password,
		Database:       database,
		InsertDate:     insertDate,
	}, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, *verrors.VError) {
	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if strings.TrimSpace(email) == "" {
		return nil, verrors.New(verrors.InvalidID, verrors.UserIDNegativeMessage)
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
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(email)
	if nil != err {
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, verrors.New(verrors.NotFound, verrors.UserNotFoundMessage)
	}

	var id int64
	var obtainedEmail string
	var profilePicture string
	var password string
	var database string
	var insertDate int64

	rows.Scan(
		&id,
		&obtainedEmail,
		&profilePicture,
		&password,
		&database,
		&insertDate,
	)

	return &models.User{
		ID:             id,
		Email:          obtainedEmail,
		ProfilePicture: profilePicture,
		Password:       password,
		Database:       database,
		InsertDate:     insertDate,
	}, nil
}

func DeleteUser(db *sql.DB, id int64) *verrors.VError {
	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if 0 >= id {
		return verrors.New(verrors.InvalidID, verrors.UserIDNegativeMessage)
	}

	statement, err := db.Prepare("DELETE FROM user WHERE id=?")
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(id)
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	if affectedRows == 0 {
		return verrors.New(verrors.DatabaseError, verrors.UserCannotDeleteMessage)
	}

	return nil
}

func UpdateUserEmail(db *sql.DB, id int64, email string) *verrors.VError {
	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if 0 >= id {
		return verrors.New(verrors.InvalidID, verrors.UserIDNegativeMessage)
	}

	err := validations.ValidateEmail(email)
	if nil != err {
		return verrors.New(verrors.InvalidEmail, err.Error())
	}

	statement, err := db.Prepare("UPDATE user SET email = ? WHERE id = ?")
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(email, id)
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	if affectedRows == 0 {
		return verrors.New(verrors.DatabaseError, verrors.UserCannotUpdateMessage)
	}

	return nil
}

func UpdateUserProfilePicture(db *sql.DB, id int64, profilePic string) *verrors.VError {
	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if 0 >= id {
		return verrors.New(verrors.InvalidID, verrors.UserIDNegativeMessage)
	}

	statement, err := db.Prepare("UPDATE user SET profile_pic = ? WHERE id = ?")

	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	res, err := statement.Exec(profilePic, id)
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	if affectedRows == 0 {
		return verrors.New(verrors.DatabaseError, verrors.UserCannotUpdateMessage)
	}

	return nil
}

func Login(
	db *sql.DB,
	serviceID string,
	registeredDomains []string,
	secret string,
	email string,
	password string,
	device *models.Device,
) (*string, *verrors.VError) {
	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if strings.TrimSpace(serviceID) == "" {
		return nil, verrors.Unexpected(verrors.ServiceIDEmptyMessage)
	}

	if 0 >= len(registeredDomains) {
		return nil, verrors.Unexpected(verrors.RegisteredDomainsEmptyMessage)
	}

	if strings.TrimSpace(secret) == "" {
		return nil, verrors.Unexpected(verrors.SecretEmptyMessage)
	}

	if strings.TrimSpace(email) == "" {
		return nil, verrors.New(verrors.InvalidEmail, verrors.EmailEmptyMessage)
	}

	if nil == device {
		return nil, verrors.New(verrors.InvalidRequest, verrors.DeviceEmptyMessage)
	}

	if strings.TrimSpace(device.Address) == "" {
		return nil, verrors.New(verrors.InvalidRequest, verrors.DeviceAddressEmptyMessage)
	}

	if strings.TrimSpace(device.UserAgent) == "" {
		return nil, verrors.New(verrors.InvalidRequest, verrors.DeviceUserAgentEmptyMessage)
	}

	statement, err := db.Prepare("SELECT id, password FROM user WHERE email = ?")
	if nil != err {
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(email)
	if nil != err {
		return nil, verrors.New(verrors.DatabaseError, err.Error())
	}

	if !rows.Next() {
		return nil, verrors.New(verrors.AccessDenied, verrors.AccessDeniedMessage)
	}

	var userID int64
	var hash string
	rows.Scan(&userID, &hash)
	rows.Close()

	err = cryptography.CompareHash(hash, password)
	if nil != err {
		return nil, verrors.New(verrors.AccessDenied, verrors.AccessDeniedMessage)
	}

	token, err := createUserToken(
		secret,
		&models.DeviceToken{
			UserAgent: device.UserAgent,
			Address:   device.Address,
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  registeredDomains,
				Issuer:    serviceID,
				Subject:   email,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		},
	)
	if nil != err {
		return nil, verrors.New(verrors.CannotGenerateAuthToken, err.Error())
	}

	device.Token = token
	updateErr := UpdateDeviceToken(db, userID, device.UserAgent, device.Address, device.Token)
	if nil != updateErr {
		if verrors.NotFound != updateErr.Code {
			return nil, updateErr
		}

		error := CreateDevice(db, userID, device)
		if nil != error {
			return nil, error
		}
	}

	return &token, nil
}

func LoginWithAuth(db *sql.DB, secret string, token string) *verrors.VError {
	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if strings.TrimSpace(secret) == "" {
		return verrors.Unexpected(verrors.SecretEmptyMessage)
	}

	if strings.TrimSpace(token) == "" {
		return verrors.New(verrors.InvalidToken, verrors.TokenEmptyMessage)
	}

	device, err := getDeviceFromUserToken(secret, token)
	if nil != err {
		return verrors.New(verrors.InvalidToken, err.Error())
	}

	user, getUserErr := GetUserByEmail(db, device.Subject)
	if nil != getUserErr {
		return getUserErr
	}

	statement, err := db.Prepare(
		"SELECT user_id, address, user_agent FROM device WHERE user_id = ? AND address = ? AND user_agent = ? AND token = ?",
	)
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(user.ID, device.Address, device.UserAgent, token)
	if nil != err {
		return verrors.New(verrors.DatabaseError, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return verrors.New(verrors.InvalidToken, verrors.TokenInvalidMessage)
	}

	return nil
}

func ValidateUserAccount(db *sql.DB, code string) *verrors.VError {
	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	return nil
}

func createUserToken(secret string, token *models.DeviceToken) (string, error) {
	token.Seed = uuid.NewString() + time.Nanosecond.String()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error parsing token: %s", err.Error())
	}

	return tokenString, nil
}

func getDeviceFromUserToken(secret string, token string) (*models.DeviceToken, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&models.DeviceToken{},
		func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
	)

	if nil != err {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*models.DeviceToken); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New(verrors.TokenInvalidMessage)
}
