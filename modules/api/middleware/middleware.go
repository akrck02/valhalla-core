package middleware

import (
	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/errors"
)

type Middleware func(context *models.ApiContext) *errors.ApiError
