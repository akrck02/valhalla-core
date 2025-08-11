package middleware

import (
	"net/http"

	"github.com/akrck02/valhalla-core/database"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

const AuthorizationHeader = "Authorization"

func Security(context *apimodels.APIContext) *verrors.APIError {
	// Check if endpoint is registered and secured
	if !context.Trazability.Endpoint.Secured {
		return nil
	}

	// Check if token is empty
	if context.Request.Authorization == "" {
		return &verrors.APIError{
			Status: http.StatusForbidden,
			VError: *verrors.Unauthorized(verrors.TokenEmptyMessage),
		}
	}

	context.Database = database.GetConnection()

	// Check if token is valid
	if !tokenIsValid(context.Request.Authorization) {
		return &verrors.APIError{
			Status: http.StatusForbidden,
			VError: *verrors.Unauthorized(verrors.TokenInvalidMessage),
		}
	}

	return nil
}

// Check if token is valid
func tokenIsValid(_ string) bool {
	return true
}
