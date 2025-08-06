package models

type User struct {
	Id             int64  `json:"id,jsonomit"`
	Email          string `json:"em,omitempty"`
	ProfilePicture string `json:"pp,omitempty"`
	Password       string `json:"-"`
	Database       string `json:"-"`
	InsertDate     int64  `json:"ins,omitempty"`
}
