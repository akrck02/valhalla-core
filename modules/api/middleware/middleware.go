// Package middleware provides functions to modify the context of a htttp request or response
package middleware

import (
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

type Middleware func(context *apimodels.ApiContext) *verrors.APIError
