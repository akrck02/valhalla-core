// Package userservice provides functions to handle the api http request related to a user
package userservice

import (
	"net/http"
	"strconv"

	"github.com/akrck02/valhalla-core/database/dal"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func Register(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	defer context.Database.Close()

	tx, err := context.Database.Begin()
	if err != nil {
		return nil, verrors.NewAPIError(verrors.New(verrors.DatabaseError, err.Error()))
	}

	defer tx.Rollback()

	user := context.Request.Body.(models.User)
	userID, registerErr := dal.RegisterUser(context.Database, &user)
	if nil != registerErr {
		return nil, verrors.NewAPIError(registerErr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: userID,
	}, nil
}

func Get(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	defer context.Database.Close()

	id := context.Request.Params["id"]
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, verrors.NewAPIError(verrors.New(verrors.InvalidRequest, err.Error()))
	}

	user, getUserErr := dal.GetUser(context.Database, userID)
	if nil != getUserErr {
		return nil, verrors.NewAPIError(getUserErr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: user,
	}, nil
}

func GetByEmail(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	defer context.Database.Close()

	email := context.Request.Params["email"]
	user, err := dal.GetUserByEmail(context.Database, email)
	if nil != err {
		return nil, verrors.NewAPIError(err)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: user,
	}, nil
}

func UpdatePassword(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	defer context.Database.Close()

	return nil, nil
}

func UpdateProfilePicture(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	return nil, nil
}

func UpdateEmail(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	return nil, nil
}

func Login(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	return nil, nil
}
