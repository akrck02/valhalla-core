package middleware

import (
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

func Checks(context *apimodels.APIContext) *verrors.APIError {
	checkError := context.Trazability.Endpoint.Checks(context)
	if checkError != nil {
		return checkError
	}

	return nil
}
