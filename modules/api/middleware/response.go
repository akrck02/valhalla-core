package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/sdk/logger"
)

const ContentTypeHeader = "Content-Type"

// EmptyResponse struct represents an empty response
type EmptyResponse struct{}

// Response handle middleware function
func Response(context *apimodels.ApiContext, writer http.ResponseWriter) {
	switch context.Trazability.Endpoint.ResponseMimeType {
	case apimodels.MimeApplicationJson:
		sendJSONHandlingErrors(context, writer)
	default:
		sendResponseCatchingErrors(context, writer)
	}
}

// SendResponse sens a HTTP response
func SendResponse(w http.ResponseWriter, status int, response interface{}, contentType apimodels.MimeType) {
	w.Header().Set(ContentTypeHeader, string(contentType))
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func sendJSONHandlingErrors(context *apimodels.ApiContext, writer http.ResponseWriter) {
	// calculate the time of the request
	start := time.Now()

	// execute the function
	response, responseError := context.Trazability.Endpoint.Listener(context)

	// calculate the time of the response
	end := time.Now()
	elapsed := end.Sub(start)

	// if something went wrong, return error
	if nil != responseError {
		logger.Error(context.Trazability.Endpoint.Path, elapsed.Nanoseconds(), "ns")
		SendResponse(writer, responseError.Status, responseError, apimodels.MimeApplicationJson)
		return
	}

	// if response is nil, return {}
	if nil == response {
		response = &apimodels.Response{
			Code:     http.StatusNoContent,
			Response: EmptyResponse{},
		}
	}

	// send response
	response.ResponseTime = elapsed.Nanoseconds()
	context.Response = *response
	SendResponse(writer, response.Code, response, apimodels.MimeApplicationJson)
	logger.Success(context.Trazability.Endpoint.Path, response.ResponseTime, "ns")
}

// Send HTTP response catching errors
func sendResponseCatchingErrors(context *apimodels.ApiContext, writer http.ResponseWriter) {
	// execute the function
	result, responseError := context.Trazability.Endpoint.Listener(context)

	// if something went wrong, return error
	if nil != responseError {
		SendResponse(writer, responseError.Status, responseError, apimodels.MimeApplicationJson)
		return
	}

	// if response is nil, return nothing
	if nil == result {
		SendResponse(writer, http.StatusNoContent, nil, context.Trazability.Endpoint.ResponseMimeType)
		return
	}

	// send response
	context.Response = *result
	SendResponse(writer, result.Code, result.Response, context.Trazability.Endpoint.ResponseMimeType)
}
