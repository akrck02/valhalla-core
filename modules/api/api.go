package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/database/tables"
	"github.com/akrck02/valhalla-core/modules/api/configuration"
	"github.com/akrck02/valhalla-core/modules/api/middleware"
	"github.com/akrck02/valhalla-core/modules/api/models"
	"github.com/akrck02/valhalla-core/modules/api/services"
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/logger"
)

const API_PATH = "/"
const CONTENT_TYPE_HEADER = "Content-Type"
const ENV_FILE_PATH = ".env"

// ApiMiddlewares is a list of middleware functions that will be applied to all API requests
// this list can be modified to add or remove middlewares
// the order of the middlewares is important, it will be applied in the order they are listed
var ApiMiddlewares = []middleware.Middleware{
	middleware.Security,
	middleware.Trazability,
	middleware.Checks,
}

func startApi(configuration configuration.APIConfiguration, endpoints []apimodels.Endpoint) {

	// show log app title and start router
	log.Println("-----------------------------------")
	log.Println(" ", configuration.ApiName, " ")
	log.Println("-----------------------------------")

	// Add API path to endpoints
	newEndpoints := []apimodels.Endpoint{}
	for _, endpoint := range endpoints {
		endpoint.Path = API_PATH + configuration.ApiName + "/" + configuration.Version + "/" + endpoint.Path
		newEndpoints = append(newEndpoints, endpoint)
	}

	// Register endpoints
	registerEndpoints(newEndpoints)

	// Start listening HTTP requests
	log.Printf("API started on http://%s:%s%s", configuration.Ip, configuration.Port, API_PATH)
	state := http.ListenAndServe(configuration.Ip+":"+configuration.Port, nil)
	log.Print(state.Error())

}

func registerEndpoints(endpoints []apimodels.Endpoint) {

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
			writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))
			writer.Header().Set("Access-Control-Allow-Methods", os.Getenv("CORS_METHODS"))
			writer.Header().Set("Access-Control-Allow-Headers", os.Getenv("CORS_HEADERS"))
			writer.Header().Set("Access-Control-Max-Age", os.Getenv("CORS_MAX_AGE"))

			// create basic api context
			context := &apimodels.ApiContext{
				Trazability: apimodels.Trazability{
					Endpoint: endpoint,
				},
			}

			// Get request data
			err := middleware.Request(reader, context)
			if nil != err {
				middleware.SendResponse(writer, err.Status, err, apimodels.MimeApplicationJson)
				return
			}

			// Apply middleware to the request
			err = applyMiddleware(context)
			if nil != err {
				middleware.SendResponse(writer, err.Status, err, apimodels.MimeApplicationJson)
				return
			}

			// Execute the endpoint
			middleware.Response(context, writer)
		})
	}
}

func setEndpointDefaults(endpoint *apimodels.Endpoint) {

	if nil == endpoint.Checks {
		endpoint.Checks = services.EmptyCheck
	}

	if nil == endpoint.Listener {
		endpoint.Listener = services.NotImplemented
	}

	if endpoint.RequestMimeType == "" {
		endpoint.RequestMimeType = apimodels.MimeApplicationJson
	}

	if endpoint.ResponseMimeType == "" {
		endpoint.ResponseMimeType = apimodels.MimeApplicationJson
	}

}

func applyMiddleware(context *apimodels.ApiContext) *errors.ApiError {

	for _, middleware := range ApiMiddlewares {
		err := middleware(context)
		if nil != err {
			return err
		}
	}

	return nil
}

func Start() {
	configuration := configuration.LoadConfiguration(ENV_FILE_PATH)

	db, err := database.Connect("valhalla.db")
	if nil != err {
		logger.Error(err)
		return
	}

	err = tables.UpdateDatabaseTablesToLatestVersion(".", tables.MainDatabase, db)
	if nil != err {
		logger.Error(err)
		return
	}

	db.Close()

	startApi(configuration, EndpointRegistry)
}
