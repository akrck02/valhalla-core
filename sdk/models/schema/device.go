package schema

type Device struct {
	Address    string `json:"address,omitempty"`
	UserAgent  string `json:"user_agent,omitempty"`
	Token      string `json:"token,omitempty"`
	InsertDate int64  `json:"insert_date,omitempty"`
	UpdateDate int64  `json:"update_date,omitempty"`
}
