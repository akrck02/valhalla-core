// Package userservice provides functions to handle the api http request related to a user
package controllers

import (
	"io"
	"net/http"
	"path"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	inout "github.com/akrck02/valhalla-core/sdk/io"
	"github.com/akrck02/valhalla-core/sdk/models"
	"github.com/akrck02/valhalla-core/services"
)

// USER REGISTER
var UserRegisterEndpoint = apimodels.Endpoint{
	Path:     "users",
	Method:   apimodels.PostMethod,
	Listener: Register,
	Secured:  false,
	Database: true,
}

func Register(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	body := context.Request.Body.(io.ReadCloser)
	user := models.User{}
	err := inout.ParseJSON(&body, &user)
	if nil != err {
		return nil, verrors.NewAPIError(verrors.New(verrors.InvalidRequestErrorCode, "Request body has invalid format."))
	}

	context.Request.Body = user

	userID, rerr := services.RegisterUser(context.Database, context.Request.Body.(models.User))

	if nil != rerr {
		return nil, verrors.NewAPIError(rerr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: userID,
	}, nil
}

var UserGetEndpoint = apimodels.Endpoint{
	Path:     "users/{id}",
	Method:   apimodels.GetMethod,
	Listener: Get,
	Secured:  true,
}

func Get(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	id, err := context.Request.GetParamInt64("id")
	if err != nil {
		return nil, verrors.NewAPIError(verrors.New(verrors.InvalidRequestErrorCode, err.Error()))
	}

	user, getUserErr := services.GetUser(context.Database, *id)
	if nil != getUserErr {
		return nil, verrors.NewAPIError(getUserErr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: user,
	}, nil
}

var UserGetByEmailEndpoint = apimodels.Endpoint{
	Path:     "users/email",
	Method:   apimodels.GetMethod,
	Listener: GetByEmail,
	Secured:  true,
}

func GetByEmail(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	email := context.Request.Params["email"]
	user, err := services.GetUserByEmail(context.Database, email)
	if nil != err {
		return nil, verrors.NewAPIError(err)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: user,
	}, nil
}

var UserUpdatePasswordEndpoint = apimodels.Endpoint{
	Path:     "users/{id}/password",
	Method:   apimodels.PatchMethod,
	Listener: UpdatePassword,
	Secured:  true,
}

func UpdatePassword(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	_, err := context.Request.GetParamInt64("id")
	if err != nil {
		return nil, verrors.NewAPIError(verrors.New(verrors.InvalidRequestErrorCode, err.Error()))
	}

	return nil, nil
}

var UserUpdateProfilePicEndpoint = apimodels.Endpoint{
	Path:            "users/{id}/profile/picture",
	Method:          apimodels.PatchMethod,
	Listener:        UpdateProfilePicture,
	IsMultipartForm: true,
	Secured:         true,
}

func UpdateProfilePicture(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	id, err := context.Request.GetParamInt64("id")
	if err != nil {
		return nil, verrors.NewAPIError(verrors.InvalidRequest(err.Error()))
	}

	fileh, ferr := context.Request.GetFile("profile_picture", 3145728)
	if nil != ferr {
		return nil, verrors.NewAPIError(verrors.InvalidRequest(ferr.Error()))
	}

	file, err := fileh.Open()
	if nil != err {
		return nil, verrors.NewAPIError(verrors.InvalidRequest(err.Error()))
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if nil != err {
		return nil, verrors.NewAPIError(verrors.InvalidRequest(err.Error()))
	}

	// Save the profile picture
	uerr := services.UpdateUserProfilePicture(context.Database, *id, &data, path.Ext(fileh.Filename))
	if nil != uerr {
		return nil, verrors.NewAPIError(uerr)
	}

	return nil, nil
}

var UserUpdateEmailEndpoint = apimodels.Endpoint{
	Path:     "users/{id}/email",
	Method:   apimodels.PatchMethod,
	Listener: UpdateEmail,
	Secured:  true,
}

func UpdateEmail(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	return nil, nil
}

var LoginEndpoint = apimodels.Endpoint{
	Path:     "users/login",
	Method:   apimodels.PostMethod,
	Listener: Login,
	Secured:  false,
	Database: true,
}

func Login(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {

	body := context.Request.Body.(io.ReadCloser)
	var params map[string]string
	inout.ParseJSON(&body, &params)

	token, err := services.Login(
		context.Database,
		"valhalla-core",
		context.Configuration.JWTRegisteredDomains,
		context.Configuration.JWTSecret,
		params["em"],
		params["ps"],
		&models.Device{
			Address:   context.Request.IP,
			UserAgent: context.Request.UserAgent,
		},
	)

	if nil != err {
		return nil, verrors.NewAPIError(err)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: token,
	}, nil
}

var LoginWithAuthEndpoint = apimodels.Endpoint{
	Path:     "users/login/auth",
	Method:   apimodels.PostMethod,
	Listener: LoginWithAuth,
	Secured:  true,
}

func LoginWithAuth(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	return &apimodels.Response{
		Code: http.StatusOK,
	}, nil
}
