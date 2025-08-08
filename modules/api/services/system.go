package services

import (
	"net/http"

	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
)

func Health(context *apimodels.ApiContext) (*apimodels.Response, *errors.ApiError) {
	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: "OK",
	}, nil
}

func EmptyCheck(context *apimodels.ApiContext) *errors.ApiError {
	return nil
}

func NotImplemented(context *apimodels.ApiContext) (*apimodels.Response, *errors.ApiError) {
	return nil, &errors.ApiError{
		Status: http.StatusNotImplemented,
		VError: errors.TODO(),
	}
}
