package middleware

import (
	"time"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

func Trazability(context *apimodels.ApiContext) *verrors.APIError {
	time := time.Now().UnixMilli()

	context.Trazability = apimodels.Trazability{
		Endpoint:  context.Trazability.Endpoint,
		Timestamp: &time,
	}

	return nil
}
