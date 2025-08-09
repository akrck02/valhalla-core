// Package userservice provides functions to handle the api http request related to a user
package userservice

import (
	"net/http"

	"github.com/akrck02/valhalla-core/database/dal"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func Register(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	defer context.Database.Close()

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
