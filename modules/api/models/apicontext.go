package apimodels

import (
	"database/sql"

	"github.com/akrck02/valhalla-core/modules/api/configuration"
)

type APIContext struct {
	Configuration configuration.APIConfiguration
	Trazability   Trazability
	Request       Request
	Response      Response
	Database      *sql.DB
}
