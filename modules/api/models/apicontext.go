package apimodels

import "database/sql"

type APIContext struct {
	Trazability Trazability
	Request     Request
	Response    Response
	Database    *sql.DB
}
