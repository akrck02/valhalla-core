package services

import (
	"net/http"

	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
)

func Health(context *models.ApiContext) (*models.Response, *errors.VError) {
	return &models.Response{
		Code:     http.StatusOK,
		Response: "OK",
	}, nil
}

func EmptyCheck(context *models.ApiContext) *errors.VError {
	return nil
}

func NotImplemented(context *models.ApiContext) (*models.Response, *errors.VError) {

	return nil, &errors.VError{
		Code:    errors.NotImplemented,
		Message: "Not implemented",
		Status:  http.StatusNotImplemented,
	}
}
