package models

type Project struct {
	Name         string `json:"nm,omitempty"`
	Description  string `json:"ds,omitempty"`
	Owner        int    `json:"ow,omitempty"`
	CreationDate int64  `json:"cd,omitempty"`
	UpdateDate   int64  `json:"ud,omitempty"`
}
