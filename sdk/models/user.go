package models

type User struct {
	ID             int64  `json:"id,omitempty"`
	Email          string `json:"em,omitempty"`
	ProfilePicture string `json:"pp,omitempty"`
	Password       string `json:"ps,omitempty"`
	Database       string `json:"-"`
	InsertDate     int64  `json:"ins,omitempty"`
}
