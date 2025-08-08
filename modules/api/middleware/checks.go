package middleware

import (
	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
)

func Checks(context *apimodels.ApiContext) *errors.ApiError {

	checkError := context.Trazability.Endpoint.Checks(context)
	if checkError != nil {
		return checkError
	}

	return nil
}
