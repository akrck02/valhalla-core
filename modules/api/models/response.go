package apimodels

type Response struct {
	Code         int         `json:"-"`
	Response     interface{} `json:"response,omitempty"`
	ResponseTime int64       `json:"response_time,omitempty"`
}
