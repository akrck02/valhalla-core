package apimodels

type Trazability struct {
	Endpoint  Endpoint `json:"endpoint"`
	Timestamp *int64   `json:"timestamp"`
}
