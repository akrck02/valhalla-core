package api

import (
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	userservice "github.com/akrck02/valhalla-core/modules/api/services/user"
)

var EndpointRegistry = []apimodels.Endpoint{
	userservice.UserRegisterEndpoint,
	userservice.UserGetEndpoint,
	userservice.UserGetByEmailEndpoint,
	userservice.UserUpdateEmailEndpoint,
	userservice.UserUpdatePasswordEndpoint,
	userservice.UserUpdateProfilePicEndpoint,
}
