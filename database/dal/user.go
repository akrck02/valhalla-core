package dal

import (
	"github.com/akrck02/valhalla-core/database"
	"github.com/akrck02/valhalla-core/errors"
	model "github.com/akrck02/valhalla-core/models"
)

func RegisterUser(user model.User, auth model.UserAuth) (*string, error) {
	return nil, errors.TODO()
}

func GetUser(id string) (*string, error) {

	_, err := database.Connect()
	if nil != err {
		return nil, err
	}

	return nil, errors.TODO()
}

func GetUserByEmail(email string) (*string, error) {
	return nil, errors.TODO()
}

func DeleteUser(id string) error {
	return errors.TODO()
}

func UpdateUser(id string, user string) error {
	return errors.TODO()
}

func UpdateUserProfilePicture(id string, picture []byte) error {
	return errors.TODO()
}

func Login(email string, device string) (*string, error) {
	return nil, errors.TODO()
}

func LoginWithAuth(id string, token string) (*string, error) {
	return nil, errors.TODO()
}

func ValidateUserAccount(code string) error {
	return errors.TODO()
}
