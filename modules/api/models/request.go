package apimodels

import "mime/multipart"

type Request struct {
	Authorization string                             `json:"authorization,omitempty"`
	Ip            string                             `json:"ip,omitempty"`
	UserAgent     string                             `json:"userAgent,omitempty"`
	Params        map[string]string                  `json:"params,omitempty"`
	Body          any                                `json:"body,omitempty"`
	Headers       map[string]string                  `json:"headers,omitempty"`
	Files         map[string][]*multipart.FileHeader `json:"files,omitempty"`
}
