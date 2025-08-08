package userservice

import (
	"net/http"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/dal"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/logger"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func Register(context *apimodels.ApiContext) (*apimodels.Response, *errors.ApiError) {
	db, err := database.Connect("valhalla.db")
	if nil != err {
		return nil, &errors.ApiError{
			Status: http.StatusInternalServerError,
			VError: errors.VError{
				Code:    errors.DatabaseError,
				Message: "Cannot connect to database.",
			},
		}
	}

	var user models.User = context.Request.Body.(models.User)
	logger.Log(user)
	userId, registerErr := dal.RegisterUser(db, &user)
	if nil == registerErr {
		return nil, &errors.ApiError{
			Status: http.StatusInternalServerError,
			VError: *registerErr,
		}
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: userId,
	}, nil
}
