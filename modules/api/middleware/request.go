package middleware

import (
	"net/http"
	"regexp"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

const UserAgentHeader = "User-Agent"

// Request handler middleware function
func Request(r *http.Request, context *apimodels.ApiContext) *verrors.APIError {
	parserError := r.ParseForm()

	if parserError != nil {
		return &verrors.APIError{
			Status: http.StatusBadRequest,
			VError: verrors.VError{
				Code:    verrors.InvalidRequest,
				Message: parserError.Error(),
			},
		}
	}

	context.Request = apimodels.Request{
		Authorization: r.Header.Get(AuthorizationHeader),
		Ip:            r.Host,
		UserAgent:     r.Header.Get(UserAgentHeader),
		Headers:       map[string]string{},
		Body:          r.Body,
		Params:        map[string]string{},
	}

	// If files are present, add them to the context
	if r.MultipartForm != nil {
		context.Request.Files = r.MultipartForm.File
	}

	// Add headers to the context
	for key, value := range r.Header {
		for _, v := range value {
			context.Request.Headers[key] = v
		}
	}

	// Add params to the context
	for key, value := range r.Form {
		for _, v := range value {
			context.Request.Params[key] = v
		}
	}

	// Get possible url path parameters
	pathParams := getPathParamNames(context.Trazability.Endpoint.Path)
	for _, param := range pathParams {
		context.Request.Params[param] = r.PathValue(param)
	}

	return nil
}

// Get the path param names
func getPathParamNames(path string) []string {
	params := []string{}

	// regex to find path parameters
	regex, err := regexp.Compile("{(.*?)}")
	if err != nil {
		return params
	}

	params = regex.FindAllString(path, -1)

	if params == nil {
		params = []string{}
	}

	for i, param := range params {
		params[i] = param[1 : len(param)-1]
	}

	return params
}
