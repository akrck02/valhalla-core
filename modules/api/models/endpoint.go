package apimodels

import (
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

type Endpoint struct {
	Path   string     `json:"path,omitempty"`
	Method HttpMethod `json:"method,omitempty"`

	RequestMimeType  MimeType `json:"requestMimeType,omitempty"`
	ResponseMimeType MimeType `json:"responseMimeType,omitempty"`

	Listener EndpointListener `json:"-"`
	Checks   EndpointCheck    `json:"-"`

	Secured  bool `json:"secured,omitempty"`
	Database bool `json:"-"`
}

type (
	EndpointCheck    func(context *ApiContext) *verrors.APIError
	EndpointListener func(context *ApiContext) (*Response, *verrors.APIError)
)
