package models

type Device struct {
	Address   string `json:"ad,omitempty"`
	UserAgent string `json:"ua,omitempty"`
	Token     string `json:"tk,omitempty"`
}
