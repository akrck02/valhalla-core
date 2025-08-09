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
			VError: verrors.VError{
				Code:    verrors.InvalidToken,
				Message: "Missing token",
			},
		}
	}

	db, err := database.Connect("valhalla.db")
	if nil != err {
		return verrors.NewAPIError(&verrors.VError{
			Code:    verrors.DatabaseError,
			Message: "Cannot connect to database.",
		})
	}

	context.Database = db

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
