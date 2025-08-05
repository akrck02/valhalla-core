package models

type Trazability struct {
	Endpoint  Endpoint `json:"endpoint"`
	Timestamp *int64   `json:"timestamp"`
}
