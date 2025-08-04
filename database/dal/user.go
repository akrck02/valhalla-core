package dal

import (
	"github.com/akrck02/valhalla-core/errors"
	model "github.com/akrck02/valhalla-core/models"
)

func RegisterUser(user model.User, auth model.UserAuth) (*string, errors.VError) {

	return nil, errors.TODO()
}

func GetUser(id string) (*string, errors.VError) {

	// _, err := database.Connect()
	// if nil != err {
	// 	return nil, err
	// }

	return nil, errors.TODO()
}

func GetUserByEmail(email string) (*string, errors.VError) {
	return nil, errors.TODO()
}

func DeleteUser(id string) errors.VError {
	return errors.TODO()
}

func UpdateUser(id string, user string) errors.VError {
	return errors.TODO()
}

func UpdateUserProfilePicture(id string, picture []byte) errors.VError {
	return errors.TODO()
}

func Login(email string, device string) (*string, errors.VError) {
	return nil, errors.TODO()
}

func LoginWithAuth(id string, token string) (*string, errors.VError) {
	return nil, errors.TODO()
}

func ValidateUserAccount(code string) errors.VError {
	return errors.TODO()
}
