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

	devices, err := dal.GetUserDevices(context.Database, *context.Request.UserID)
	if nil != err {
		return nil, verrors.NewAPIError(err)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: devices,
	}, nil
}

var GetDeviceEndpoint = apimodels.Endpoint{
	Path:     "devices",
	Method:   apimodels.GetMethod,
	Listener: GetDevice,
	Secured:  true,
}

func GetDevice(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {

	userID, err := context.Request.GetParamInt64("userId")
	if nil != err {
		return nil, verrors.NewAPIError(verrors.InvalidRequest(err.Error()))
	}

	devices, verr := dal.GetDevice(context.Database, *userID, context.Request.Params["user_agent"], context.Request.Params["address"])
	if nil != verr {
		return nil, verrors.NewAPIError(verr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: devices,
	}, nil
}
