// Package services provides functions to handle http requests in the api
package services

import (
	"net/http"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

func Health(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: "OK",
	}, nil
}

func EmptyCheck(context *apimodels.APIContext) *verrors.APIError {
	return nil
}

func NotImplemented(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	return nil, &verrors.APIError{
		Status: http.StatusNotImplemented,
		VError: verrors.TODO(),
	}
}
