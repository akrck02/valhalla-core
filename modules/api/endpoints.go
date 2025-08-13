package api

import (
	"github.com/akrck02/valhalla-core/modules/api/controllers"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
)

var EndpointRegistry = []apimodels.Endpoint{
	controllers.UserRegisterEndpoint,
	controllers.UserGetEndpoint,
	controllers.UserGetByEmailEndpoint,
	controllers.UserUpdateEmailEndpoint,
	controllers.UserUpdatePasswordEndpoint,
	controllers.UserUpdateProfilePicEndpoint,
	controllers.LoginEndpoint,
	controllers.LoginWithAuthEndpoint,
	controllers.GetUserDevicesEndpoint,
}
