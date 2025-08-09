package apimodels

import "database/sql"

type ApiContext struct {
	Trazability Trazability
	Request     Request
	Response    Response
	Database    *sql.DB
}
