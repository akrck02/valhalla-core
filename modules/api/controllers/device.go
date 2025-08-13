package controllers

import (
	"net/http"

	"github.com/akrck02/valhalla-core/database/dal"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

var GetUserDevicesEndpoint = apimodels.Endpoint{
	Path:     "devices/user",
	Method:   apimodels.GetMethod,
	Listener: GetUserDevices,
	Secured:  true,
}

func GetUserDevices(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {

	devices, err := dal.GetDevice(context.Database, *context.Request.UserID, context.Request.UserAgent, context.Request.IP)
	if nil != err {
		return nil, verrors.NewAPIError(err)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: devices,
	}, nil
}
