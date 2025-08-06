package middleware

import (
	"log"
	"net/http"

	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
)

const AUTHORITATION_HEADER = "Authorization"

func Security(context *models.ApiContext) *errors.ApiError {

	// Check if endpoint is registered and secured
	if !context.Trazability.Endpoint.Secured {
		return nil
	}

	log.Printf("Endpoint %s is secured", context.Trazability.Endpoint.Path)

	// Check if token is empty
	if context.Request.Authorization == "" {
		return &errors.ApiError{
			Status: http.StatusForbidden,
			Error: errors.VError{
				Code:    errors.InvalidToken,
				Message: "Missing token",
			},
		}
	}

	// Check if token is valid
	if !tokenIsValid(context.Request.Authorization) {
		return &errors.ApiError{
			Status: http.StatusForbidden,
			Error: errors.VError{
				Code:    errors.InvalidToken,
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
