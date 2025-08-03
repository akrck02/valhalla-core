package model

type User struct {
	Email          string `json:"em"`
	Username       string `json:"us"`
	ProfilePicture string `json:"pp"`
}

type UserAuth struct {
	Password string `json:"ps"`
	Token    string `json:"tk"`
}
