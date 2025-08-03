package data

import (
	"github.com/akrck02/valhalla-core/data/connection"
	model "github.com/akrck02/valhalla-core/models"
)

func register(user model.User, auth model.UserAuth) (*string, error) {
	return nil, nil
}

func get(id string) (*string, error) {

	_, err := connection.GetSqlite()
	if nil != err {
		return nil, err
	}

	return nil, nil
}

func getByEmail(email string) (*string, error) {
	return nil, nil
}

func delete(id string) error {
	return nil
}

func update(id string, user string) error {
	return nil
}

func updateProfilePicture(id string, picture []byte) error {
	return nil
}

func login(email string, device string) (*string, error) {
	return nil, nil
}

func loginWithAuth(id string, token string) (*string, error) {
	return nil, nil
}

func validateAccount(code string) error {
	return nil
}
