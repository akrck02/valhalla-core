package model

type Project struct {
	Name         string `json:"nm"`
	Description  string `json:"ds"`
	Owner        int    `json:"ow"`
	CreationDate int64  `json:"cd"`
	UpdateDate   int64  `json:"ud"`
}
