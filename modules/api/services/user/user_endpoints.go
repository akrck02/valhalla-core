package userservice

import apimodels "github.com/akrck02/valhalla-core/modules/api/models"

var UserRegisterEndpoint = apimodels.Endpoint{
	Path:     "users",
	Method:   apimodels.PostMethod,
	Listener: Register,
	Checks:   RegisterCheck,
	Secured:  false,
	Database: true,
}

var UserGetEndpoint = apimodels.Endpoint{
	Path:     "users/{id}",
	Method:   apimodels.GetMethod,
	Listener: Get,
	// Checks:   RegisterCheck,
	Secured: true,
}

var UserGetByEmailEndpoint = apimodels.Endpoint{
	Path:     "users/email",
	Method:   apimodels.GetMethod,
	Listener: GetByEmail,
	// Checks:   RegisterCheck,
	Secured: true,
}

var UserUpdatePasswordEndpoint = apimodels.Endpoint{
	Path:     "users/{id}/password",
	Method:   apimodels.PatchMethod,
	Listener: UpdatePassword,
	// Checks:   RegisterCheck,
	Secured: true,
}

var UserUpdateProfilePicEndpoint = apimodels.Endpoint{
	Path:     "users/{id}/profile/picture",
	Method:   apimodels.PatchMethod,
	Listener: UpdateProfilePicture,
	// Checks:   RegisterCheck,
	IsMultipartForm: true,
	Secured:         true,
}

var UserUpdateEmailEndpoint = apimodels.Endpoint{
	Path:     "users/{id}/email",
	Method:   apimodels.PatchMethod,
	Listener: UpdateEmail,
	// Checks:   RegisterCheck,
	Secured: true,
}

var LoginEndpoint = apimodels.Endpoint{
	Path:     "users/login",
	Method:   apimodels.PatchMethod,
	Listener: Login,
	// Checks:   RegisterCheck,
	Secured: false,
}
