// Package api provides the valhalla http endpoints
package api

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/tables"
	"github.com/akrck02/valhalla-core/modules/api/configuration"
	"github.com/akrck02/valhalla-core/modules/api/controllers"
	"github.com/akrck02/valhalla-core/modules/api/middleware"
	apimodels "github.com/akrck02/valhalla-core/modules/api/models"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/logger"
)

const (
	APIPath           = "/"
	ContentTypeHeader = "Content-Type"
	EnvFilePath       = ".env"
)

// APIMiddleware is a list of middleware functions that will be applied to all API requests
// this list can be modified to add or remove middlewares
// the order of the middlewares is important, it will be applied in the order they are listed
var APIMiddleware = []middleware.Middleware{
	middleware.Security,
	middleware.Trazability,
	middleware.Database,
}

func startAPI(configuration *configuration.APIConfiguration, endpoints *[]apimodels.Endpoint) {
	// show log app title and start router
	log.Println("-----------------------------------")
	log.Println(" ", configuration.ApiName, " ")
	log.Println("-----------------------------------")

	// Add API path to endpoints
	newEndpoints := []apimodels.Endpoint{}
	for _, endpoint := range *endpoints {
		endpoint.Path = APIPath + configuration.ApiName + "/" + configuration.Version + "/" + endpoint.Path
		newEndpoints = append(newEndpoints, endpoint)
	}

	// Register endpoints
	registerEndpoints(newEndpoints, configuration)

	// Start listening HTTP requests
	log.Printf("API started on http://%s:%s%s", configuration.Ip, configuration.Port, APIPath)
	state := http.ListenAndServe(configuration.Ip+":"+configuration.Port, nil)
	log.Print(state.Error())
}

func registerEndpoints(endpoints []apimodels.Endpoint, configuration *configuration.APIConfiguration) {
	for _, endpoint := range endpoints {

		switch endpoint.Method {
		case apimodels.GetMethod:
			endpoint.Path = fmt.Sprintf("GET %s", endpoint.Path)
		case apimodels.PostMethod:
			endpoint.Path = fmt.Sprintf("POST %s", endpoint.Path)
		case apimodels.PutMethod:
			endpoint.Path = fmt.Sprintf("PUT %s", endpoint.Path)
		case apimodels.DeleteMethod:
			endpoint.Path = fmt.Sprintf("DELETE %s", endpoint.Path)
		case apimodels.PatchMethod:
			endpoint.Path = fmt.Sprintf("PATCH %s", endpoint.Path)
		}

		log.Printf("Endpoint %s registered. \n", endpoint.Path)

		// set defaults
		setEndpointDefaults(&endpoint)

		http.HandleFunc(endpoint.Path, func(writer http.ResponseWriter, reader *http.Request) {

			// enable CORS
			writer.Header().Set("Access-Control-Allow-Origin", configuration.CorsAccessControlAllowOrigin)
			writer.Header().Set("Access-Control-Allow-Methods", configuration.CorsAccessControlAllowMethods)
			writer.Header().Set("Access-Control-Allow-Headers", configuration.CorsAccessControlAllowHeaders)
			writer.Header().Set("Access-Control-Max-Age", configuration.CorsAccessControlMaxAge)

			// calculate the time of the request
			start := time.Now()

			// create basic api context
			context := &apimodels.APIContext{
				Trazability: apimodels.Trazability{
					Endpoint: endpoint,
				},
			}

			// Get request data
			err := middleware.Request(reader, context)
			if nil != err {
				logger.Error(
					context.Trazability.Endpoint.Path,
					time.Since(start).Microseconds(),
					"μs -",
					fmt.Sprintf("[%d]", err.Status),
					err.Message,
				)

				middleware.SendResponse(writer, err.Status, err, apimodels.MimeApplicationJSON)
				return
			}

			// Apply middleware to the request
			err = applyMiddleware(context)
			defer database.ReturnConnection(context.Database)
			if nil != err {
				logger.Error(
					context.Trazability.Endpoint.Path,
					time.Since(start).Microseconds(),
					"μs -",
					fmt.Sprintf("[%d]", err.Status),
					err.Message,
				)
				middleware.SendResponse(writer, err.Status, err, apimodels.MimeApplicationJSON)
				return
			}

			// Execute the endpoint
			middleware.Response(context, writer)
		})
	}
}

func setEndpointDefaults(endpoint *apimodels.Endpoint) {
	if nil == endpoint.Listener {
		endpoint.Listener = controllers.NotImplemented
	}

	if endpoint.RequestMimeType == "" {
		endpoint.RequestMimeType = apimodels.MimeApplicationJSON
	}

	if endpoint.ResponseMimeType == "" {
		endpoint.ResponseMimeType = apimodels.MimeApplicationJSON
	}
}

func applyMiddleware(context *apimodels.APIContext) *verrors.APIError {
	for _, middleware := range APIMiddleware {
		err := middleware(context)
		if nil != err {
			return err
		}
	}

	return nil
}

func Start() {
	configuration := configuration.LoadConfiguration(EnvFilePath)
	err := database.StartConnectionPool()
	if nil != err {
		logger.Error(err)
		return
	}

	db := database.GetConnection()
	err = tables.UpdateDatabaseTablesToLatestVersion(".", tables.MainDatabase, db)
	if nil != err {
		logger.Error(err)
		return
	}
	database.ReturnConnection(db)

	startAPI(&configuration, &EndpointRegistry)
}
