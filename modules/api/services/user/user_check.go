package userservice

import (
	"io"
	"net/http"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
	inout "github.com/akrck02/valhalla-core/sdk/io"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func RegisterCheck(context *apimodels.ApiContext) *errors.ApiError {

	var body = context.Request.Body.(io.ReadCloser)
	var user = models.User{}
	err := inout.ParseJson(&body, &user)
	if nil != err {
		return &errors.ApiError{
			Status: http.StatusInternalServerError,
			VError: *errors.New(errors.InvalidRequest, "Request body has invalid format."),
		}
	}

	context.Request.Body = user
	return nil
}
