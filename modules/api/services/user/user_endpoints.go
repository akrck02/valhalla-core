package userservice

import "github.com/akrck02/valhalla-core/modules/api/models"

var UserRegisterEndpoint = apimodels.Endpoint{
	Path:     "users",
	Method:   apimodels.GetMethod,
	Listener: Register,
	Checks:   RegisterCheck,
	Secured:  false,
	Database: false,
}
