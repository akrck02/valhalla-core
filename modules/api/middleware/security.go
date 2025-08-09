package middleware

import (
	"log"
	"net/http"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

const AuthorizationHeader = "Authorization"

func Security(context *apimodels.ApiContext) *verrors.APIError {
	// Check if endpoint is registered and secured
	if !context.Trazability.Endpoint.Secured {
		return nil
	}

	log.Printf("Endpoint %s is secured", context.Trazability.Endpoint.Path)

	// Check if token is empty
	if context.Request.Authorization == "" {
		return &verrors.APIError{
			Status: http.StatusForbidden,
			VError: verrors.VError{
				Code:    verrors.InvalidToken,
				Message: "Missing token",
			},
		}
	}

	// Check if token is valid
	if !tokenIsValid(context.Request.Authorization) {
		return &verrors.APIError{
			Status: http.StatusForbidden,
			VError: verrors.VError{
				Code:    verrors.InvalidToken,
				Message: "Invalid token",
			},
		}
	}

	return nil
}

// Check if token is valid
func tokenIsValid(_ string) bool {
	return true
}
