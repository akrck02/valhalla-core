package middleware

import (
	"time"

	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
)

func Trazability(context *models.ApiContext) *errors.ApiError {

	time := time.Now().UnixMilli()

	context.Trazability = models.Trazability{
		Endpoint:  context.Trazability.Endpoint,
		Timestamp: &time,
	}

	return nil
}
