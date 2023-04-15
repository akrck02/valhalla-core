package services

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/lang"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type validatePasswordResult struct {
	Response error.User
	Message  string
}

const MINIMUM_CHARACTERS_FOR_PASSWORD = 16

var SPECIAL_CHARATERS = []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", "|", ";", ":", "'", ",", ".", "<", ">", "?", "/", "`", "~"}

// Register HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func RegisterHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params RegisterParams
	err := c.ShouldBindJSON(&params)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_NOT_ACCEPTABLE,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var user models.User
	user.Username = params.Username
	user.Password = params.Password
	user.Email = params.Email

	var error = Register(conn, client, &user)
	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	// send response
	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User created"},
	)
}

// Register user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | *models.User: user to register
//
// [return] *models.Error: error if any
func Register(conn context.Context, client *mongo.Client, user *models.User) *models.Error {

	var checkedPass = validatePassword(user.Password)

	if checkedPass.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(checkedPass.Response),
			Message: checkedPass.Message,
		}
	}

	coll := client.Database("valhalla").Collection("user")
	found := mailExists(user.Email, conn, coll)

	if found.Email != "" {

		return &models.Error{
			Code:    utils.HTTP_STATUS_CONFLICT,
			Error:   int(error.USER_ALREADY_EXISTS),
			Message: "User already exists",
		}
	}

	user.Password = utils.EncryptSha256(user.Password)

	// register user on database
	_, err := coll.InsertOne(conn, user)

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_ALREADY_EXISTS),
			Message: "User already exists",
		}
	}

	return nil
}

// Login HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func LoginHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	json.Unmarshal([]byte(jsonData), &user)

	ip := c.ClientIP()
	address := c.Request.Header.Get("User-Agent")
	token, error := Login(conn, client, user, ip, address)

	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"code": utils.HTTP_STATUS_OK, "message": "User found", "auth": token},
	)

}

// Login user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to login
// [param] ip | string: ip address of the user
// [param] address | string: user agent of the user
//
// [return] string: auth token --> *models.Error: error if any
func Login(conn context.Context, client *mongo.Client, user models.User, ip string, address string) (string, *models.Error) {

	coll := client.Database("valhalla").Collection("user")
	found := authorizationOk(user.Username, user.Password, conn, coll)

	if found.Email == "" {
		return "", &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Message: "Forbidden",
		}
	}

	device := models.Device{Address: ip, UserAgent: address}
	token, err := AddUserDevice(conn, client, found, device)

	if err != nil {
		return "", &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Message: "Cannot generate your auth token",
		}
	}

	return token, nil
}

// Edit user HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func EditUserHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	json.Unmarshal([]byte(jsonData), &user)

	updateErr := EditUser(conn, client, user)
	if updateErr != nil {
		utils.SendResponse(c,
			updateErr.Code,
			gin.H{"http-code": updateErr.Code, "internal-code": updateErr.Error, "message": updateErr.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User updated"},
	)

}

// Edit user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to edit
//
// [return] *models.Error: error if any
func EditUser(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	users := client.Database("valhalla").Collection("user")

	// update user on database
	res, err := users.UpdateOne(conn, bson.M{"email": user.Email}, bson.M{"$set": bson.M{"username": user.Username, "password": user.Password}})

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_UPDATED),
			Message: "User not updated",
		}
	}

	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.USER_NOT_FOUND),
			Message: "Users not found",
		}
	}

	return nil
}

func DeleteUserHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	json.Unmarshal([]byte(jsonData), &user)

	deleteErr := DeleteUser(conn, client, user)
	if deleteErr != nil {
		utils.SendResponse(c,
			deleteErr.Code,
			gin.H{"http-code": deleteErr.Code, "internal-code": deleteErr.Error, "message": deleteErr.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User deleted"},
	)

}

// Delete user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to delete
//
// [return] *models.Error: error if any
func DeleteUser(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	users := client.Database("valhalla").Collection("user")

	// delete user on database
	deleteResult, err := users.DeleteOne(conn, bson.M{"email": user.Email})

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_DELETED),
			Message: "User not deleted",
		}
	}

	if deleteResult.DeletedCount == 0 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.USER_NOT_FOUND),
			Message: "User not found",
		}
	}

	return nil
}

// Check if the given password is valid
// following the next rules:
//
//		[-] At least 16 characters
//		[-] At least one special character
//		[-] At least one number
//
//	 [param] password : string: password to check
//
//	 [return] the password is valid or not
func validatePassword(password string) validatePasswordResult {

	if len(password) < MINIMUM_CHARACTERS_FOR_PASSWORD {
		return validatePasswordResult{
			Response: error.SHORT_PASSWORD,
			Message:  "Password must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_PASSWORD) + " characters",
		}
	}

	if !utils.ContainsAny(password, SPECIAL_CHARATERS) {
		return validatePasswordResult{
			Response: error.NO_SPECIAL_CHARACTERS_PASSWORD,
			Message:  "Password must have at least one special character",
		}
	}

	if !utils.IsLowerCase(password) {
		return validatePasswordResult{
			Response: error.NO_UPPER_LOWER_PASSWORD,
			Message:  "Password must have at least one uppercase character",
		}
	}

	if !utils.IsUpperCase(password) {
		return validatePasswordResult{
			Response: error.NO_UPPER_LOWER_PASSWORD,
			Message:  "Password must have at least one lowercase character",
		}
	}

	return validatePasswordResult{
		Response: 200,
		Message:  "Ok.",
	}
}

// Check email on database
//
//	[param] email | string The email to check
//	[param] conn | context.Context The connection to the database
//
//	[return] model.User : The user found or empty
func mailExists(email string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "email", Value: email}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result

}

// Get if the given credentials are valid
//
//	[param] username | string : The username to check
//	[param] password | string : The password to check
//	[param] conn | context.Context : The connection to the database
//
//	[return] model.User : The user found or empty
func authorizationOk(username string, password string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "username", Value: username}, {Key: "password", Value: utils.EncryptSha256(password)}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result

}
