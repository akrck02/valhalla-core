package models

import "github.com/golang-jwt/jwt/v5"

type Device struct {
	Address    string `json:"ad,omitempty"`
	UserAgent  string `json:"ua,omitempty"`
	Token      string `json:"tk,omitempty"`
	InsertDate int64  `json:"ind,omitempty"`
	UpdateDate int64  `json:"ud,omitempty"`
}

type DeviceToken struct {
	Address   string `json:"ad,omitempty"`
	UserAgent string `json:"ua,omitempty"`
	jwt.RegisteredClaims
}
