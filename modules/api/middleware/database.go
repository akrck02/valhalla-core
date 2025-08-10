package middleware

import (
	"github.com/akrck02/valhalla-core/database"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

func Database(context *apimodels.APIContext) *verrors.APIError {
	if !context.Trazability.Endpoint.Database || nil != context.Database {
		return nil
	}

	db, err := database.Connect("valhalla.db")
	if nil != err {
		return verrors.NewAPIError(&verrors.VError{
			Code:    verrors.DatabaseErrorCode,
			Message: "Cannot connect to database.",
		})
	}

	context.Database = db

	return nil
}
