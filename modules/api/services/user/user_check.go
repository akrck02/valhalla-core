package userservice

import (
	"io"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	inout "github.com/akrck02/valhalla-core/sdk/io"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func RegisterCheck(context *apimodels.APIContext) *verrors.APIError {
	body := context.Request.Body.(io.ReadCloser)
	user := models.User{}
	err := inout.ParseJson(&body, &user)
	if nil != err {
		return verrors.NewAPIError(verrors.New(verrors.InvalidRequest, "Request body has invalid format."))
	}

	context.Request.Body = user
	return nil
}
