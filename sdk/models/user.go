package models

type User struct {
	Id             int64  `json:"id,jsonomit"`
	Email          string `json:"em,omitempty"`
	ProfilePicture string `json:"pp,omitempty"`
	InsertDate     int64  `json:"ins,omitempty"`
}

type UserAuth struct {
	UserId   int64  `json:"uid,omitempty"`
	Password string `json:"ps,omitempty"`
	Token    string `json:"tk,omitempty"`
}
