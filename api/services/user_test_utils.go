package services

import (
	"context"
	"testing"

	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/middleware"
	"github.com/akrck02/valhalla-core/mock"
	"github.com/akrck02/valhalla-core/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterMockTestUser registers a fake user
// and returns the user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
//
// [return] *models.User : user object
func RegisterMockTestUser(t *testing.T, conn context.Context, client *mongo.Client) *models.User {

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	return RegisterTestUser(t, conn, client, user)

}

// RegisterTestUser registers a user and handles the error
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
//
// [return] *models.User : user object
func RegisterTestUser(t *testing.T, conn context.Context, client *mongo.Client, user *models.User) *models.User {

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return nil
	}

	return user
}

// RegisterTestUser registers a fake user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
// [param] status | int : HTTP status
// [param] errorcode | int : error code
func RegisterTestUserWithError(t *testing.T, conn context.Context, client *mongo.Client, user *models.User, status int, errorcode int) {

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered.")
		return
	}

	if err.Status != status || err.Error != errorcode {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)

}

// LoginTestUser logs in a fake user
// and returns the token
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
//
// [return] string : token
func LoginTestUser(t *testing.T, conn context.Context, client *mongo.Client, user *models.User, ip string, userAgent string) string {

	log.FormattedInfo("Logging in user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)

	token, err := Login(conn, client, user, ip, userAgent)

	if err != nil {
		t.Error("The user was not logged in", err)
		return ""
	}

	if token == "" {
		t.Error("The token is empty")
		return ""
	}

	_, err = middleware.IsTokenValid(client, token)

	if err != nil {
		t.Error("The token was not validated", err)
		return ""
	}

	log.Info("User logged in")
	log.FormattedInfo("Token: ${0}", token)
	return token

}

// LoginTestUserWithError logs in a fake user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
// [param] ip | string : ip address
// [param] userAgent | string : user agent
// [param] status | int : HTTP status
// [param] errorcode | int : error code
func LoginTestUserWithError(t *testing.T, conn context.Context, client *mongo.Client, user *models.User, ip string, userAgent string, status int, errorcode int) {

	log.FormattedInfo("Logging in user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)

	_, err := Login(conn, client, user, ip, userAgent)

	if err == nil {
		t.Error("The user was logged in.")
		return
	}

	if err.Status != status || err.Error != errorcode {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not logged in")
	log.FormattedInfo("Error: ${0}", err.Message)
}

// LoginAuth logs in a user with a token
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] email | string : email
// [param] token | string : token
func LoginAuthTestUser(t *testing.T, conn context.Context, client *mongo.Client, email string, token string) {

	log.FormattedInfo("Authenticating token: ${0}", token)

	log.Info("Checking if the token is valid")

	var authLogin = &models.AuthLogin{
		Email:     email,
		AuthToken: token,
	}

	err := LoginAuth(conn, client, authLogin, mock.Ip(), mock.Platform())

	if err != nil {
		t.Error("The token is not valid", err)
		return
	}

	log.Info("Token is valid")

}

// DeleteUser deletes a user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
func DeleteTestUser(t *testing.T, conn context.Context, client *mongo.Client, user *models.User) {

	log.FormattedInfo("Deleting user: ${0}", user.Email)

	err := DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")

}

// DeleteTestUserWithError deletes a user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
// [param] status | int : HTTP status
// [param] errorcode | int : error code
func DeleteTestUserWithError(t *testing.T, conn context.Context, client *mongo.Client, user *models.User, status int, errorcode int) {

	log.FormattedInfo("Deleting user: ${0}", user.Email)

	err := DeleteUser(conn, client, user)

	if err == nil {
		t.Error("The user was deleted.")
		return
	}

	if err.Status != status || err.Error != errorcode {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not deleted")
	log.FormattedInfo("Error: ${0}", err.Message)

}

// EditTestUser edits a user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object

func EditTestUserEmail(t *testing.T, conn context.Context, client *mongo.Client, emailChangeRequest *EmailChangeRequest) {

	log.FormattedInfo("Editing user mail: ${0}", emailChangeRequest.Email)
	log.FormattedInfo("New email: ${0}", emailChangeRequest.NewEmail)

	err := EditUserEmail(conn, client, emailChangeRequest)

	if err != nil {
		t.Error("The user was not edited", err)
		return
	}

	log.Info("User edited")

}

// EditTestUser edits a user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
// [param] status | int : HTTP status
// [param] errorcode | int : error code
func EditTestUserEmailWithError(t *testing.T, conn context.Context, client *mongo.Client, emailChangeRequest *EmailChangeRequest, status int, errorcode int) {

	log.FormattedInfo("Editing user mail: ${0}", emailChangeRequest.Email)
	log.FormattedInfo("New email: ${0}", emailChangeRequest.NewEmail)

	err := EditUserEmail(conn, client, emailChangeRequest)

	if err == nil {
		t.Error("The user was edited.")
		return
	}

	if err.Status != status || err.Error != errorcode {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not edited")
	log.FormattedInfo("Error: ${0}", err.Message)

}

// EditTestUser edits a user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
// [param] status | int : HTTP status
// [param] errorcode | int : error code
func EditTestUser(t *testing.T, conn context.Context, client *mongo.Client, user *models.User) {

	log.FormattedInfo("Editing user: ${0}", user.Email)
	err := EditUser(conn, client, user)

	if err != nil {
		t.Error("The user was not edited", err)
		return
	}

	log.Info("User edited")

}

// EditTestUser edits a user
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] user | *models.User : user object
// [param] status | int : HTTP status
// [param] errorcode | int : error code
func EditTestUserWithError(t *testing.T, conn context.Context, client *mongo.Client, user *models.User, status int, errorcode int) {

	log.FormattedInfo("Editing user: ${0}", user.Email)
	err := EditUser(conn, client, user)

	if err == nil {
		t.Error("The user was edited.")
		return
	}

	if err.Status != status || err.Error != errorcode {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not edited")
	log.FormattedInfo("Error: ${0}", err.Message)

}

// ValidateTestToken validates a token and handles the error
//
// [param] t | *testing.T : testing object
// [param] conn | context.Context : context object
// [param] client | *mongo.Client : mongo client object
// [param] token | string : token
func ValidateTestTokenWithError(t *testing.T, conn context.Context, client *mongo.Client, token string, status int, errorcode int) {

	log.FormattedInfo("Validating token: ${0}", token)

	_, err := middleware.IsTokenValid(client, token)

	if err == nil {
		t.Error("The token was validated.")
		return
	}

	if err.Status != status || err.Error != errorcode {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("Token not validated")
	log.FormattedInfo("Token not validated, error: ${0}", err.Message)

}
