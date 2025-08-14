package schema

type User struct {
	ID             int64  `json:"id,omitempty"`
	Email          string `json:"email,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Password       string `json:"password,omitempty"`
	Database       string `json:"database,omitempty"`
	InsertDate     int64  `json:"insert_date,omitempty"`
}
