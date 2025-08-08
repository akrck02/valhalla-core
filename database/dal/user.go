package dal

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/akrck02/valhalla-core/sdk/cryptography"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/akrck02/valhalla-core/sdk/validations"
	"github.com/golang-jwt/jwt/v5"
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

	if !rows.Next() {
		return nil, errors.New(errors.NotFound, "User not found.")
	}

	var obtainedId int64
	var email string
	var profilePicture string
	var password string
	var database string
	var insertDate int64

	rows.Scan(
		&obtainedId,
		&email,
		&profilePicture,
		&password,
		&database,
		&insertDate,
	)

	return &models.User{
		Id:             obtainedId,
		Email:          email,
		ProfilePicture: profilePicture,
		Password:       password,
		Database:       database,
		InsertDate:     insertDate,
	}, nil
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

	if !rows.Next() {
		return nil, errors.New(errors.NotFound, "User not found.")
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
		Id:             id,
		Email:          obtainedEmail,
		ProfilePicture: profilePicture,
		Password:       password,
		Database:       database,
		InsertDate:     insertDate,
	}, nil
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

func Login(db *sql.DB, serviceId string, registeredDomains []string, secret string, email string, password string, device *models.Device) (*string, *errors.VError) {

	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	if "" == strings.TrimSpace(serviceId) {
		return nil, errors.Unexpected("Service id cannot be empty.")
	}

	if 0 >= len(registeredDomains) {
		return nil, errors.Unexpected("Registered domains cannot be empty.")
	}

	if "" == strings.TrimSpace(secret) {
		return nil, errors.Unexpected("Secret cannot be empty.")
	}

	if "" == strings.TrimSpace(email) {
		return nil, errors.New(errors.InvalidEmail, "Email cannot be empty.")
	}

	if nil == device {
		return nil, errors.New(errors.InvalidRequest, "Device cannot be empty.")
	}

	if "" == strings.TrimSpace(device.Address) {
		return nil, errors.New(errors.InvalidRequest, "Device address cannot be empty.")
	}

	if "" == strings.TrimSpace(device.UserAgent) {
		return nil, errors.New(errors.InvalidRequest, "Device user agent cannot be empty.")
	}

	statement, err := db.Prepare("SELECT id, password FROM user WHERE email = ?")
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	rows, err := statement.Query(email)
	if nil != err {
		return nil, errors.New(errors.DatabaseError, err.Error())
	}

	if false == rows.Next() {
		return nil, errors.New(errors.AccessDenied, "Access denied.")
	}

	var userId int64
	var hash string
	rows.Scan(&userId, &hash)
	rows.Close()

	err = cryptography.CompareHash(hash, password)
	if nil != err {
		return nil, errors.New(errors.AccessDenied, "Access denied.")
	}

	token, err := createUserToken(
		secret,
		&models.DeviceToken{
			UserAgent: device.UserAgent,
			Address:   device.Address,
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  registeredDomains,
				Issuer:    serviceId,
				Subject:   email,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		},
	)
	if nil != err {
		return nil, errors.New(errors.CannotGenerateAuthToken, err.Error())
	}

	device.Token = token
	error := CreateDevice(db, userId, device)
	if nil != error {
		return nil, error
	}

	return &token, nil
}

func LoginWithAuth(db *sql.DB, token string) (*string, *errors.VError) {

	if nil == db {
		return nil, errors.Unexpected("Database connection cannot be empty.")
	}

	if "" == strings.TrimSpace(token) {
		return nil, errors.New(errors.InvalidRequest, "Token cannot be empty.")
	}

	return nil, nil
}

func ValidateUserAccount(db *sql.DB, code string) *errors.VError {

	if nil == db {
		return errors.Unexpected("Database connection cannot be empty.")
	}

	return nil
}

func createUserToken(secret string, token *models.DeviceToken) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("Error parsing token: %s", err.Error())
	}

	return tokenString, nil
}

func getDeviceFromUserToken(secret string, token string) (*models.DeviceToken, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &models.DeviceToken{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if nil != err {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*models.DeviceToken); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Token %s is invalid", token)

}
