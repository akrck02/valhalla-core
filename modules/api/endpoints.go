package api

import (
	"github.com/akrck02/valhalla-core/modules/api/controllers"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

var EndpointRegistry = []apimodels.Endpoint{
	controllers.UserRegisterEndpoint,
	controllers.UserGetEndpoint,
	controllers.UserGetByEmailEndpoint,
	controllers.UserUpdateEmailEndpoint,
	controllers.UserUpdatePasswordEndpoint,
	controllers.UserUpdateProfilePicEndpoint,
	{
		Path:     "users/ddos",
		Method:   apimodels.GetMethod,
		Secured:  true,
		Database: true,
		Listener: func(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
			return nil, nil
		},
	},
}
