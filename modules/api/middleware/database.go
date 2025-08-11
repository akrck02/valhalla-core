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

	context.Database = database.GetConnection()
	if nil == context.Database {
		return verrors.NewAPIError(&verrors.VError{
			Code:    verrors.DatabaseErrorCode,
			Message: "Cannot connect to database.",
		})
	}

	return nil
}
