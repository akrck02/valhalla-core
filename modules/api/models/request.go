package models

type Request struct {
	Authorization string            `json:"authorization,omitempty"`
	Ip            string            `json:"ip,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Params        map[string]string `json:"params,omitempty"`
	Body          interface{}       `json:"body,omitempty"`
	Headers       map[string]string `json:"headers,omitempty"`
	Files         interface{}       `json:"files,omitempty"`
}
