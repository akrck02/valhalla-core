package services

import (
	"net/http"

	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
)

func Health(context *models.ApiContext) (*models.Response, *errors.ApiError) {
	return &models.Response{
		Code:     http.StatusOK,
		Response: "OK",
	}, nil
}

func EmptyCheck(context *models.ApiContext) *errors.ApiError {
	return nil
}

func NotImplemented(context *models.ApiContext) (*models.Response, *errors.ApiError) {
	return nil, &errors.ApiError{
		Status: http.StatusNotImplemented,
		Error:  errors.TODO(),
	}
}
