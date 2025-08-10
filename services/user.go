package services

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/akrck02/valhalla-core/database/dal"
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	inout "github.com/akrck02/valhalla-core/sdk/io"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func RegisterUser(db *sql.DB, user models.User) (*int64, *verrors.VError) {
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, verrors.DatabaseError(err.Error())
	}

	defer tx.Rollback()

	userID, rerr := dal.RegisterUser(db, &user)
	if nil != rerr {
		return nil, rerr
	}

	return userID, nil
}

func GetUser(db *sql.DB, id int64) (*models.User, *verrors.VError) {
	defer db.Close()

	user, err := dal.GetUser(db, id)
	if nil != err {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, *verrors.VError) {
	defer db.Close()

	user, err := dal.GetUserByEmail(db, email)
	if nil != err {
		return nil, err
	}

	return user, nil
}

func UpdateUserProfilePicture(db *sql.DB, userID int64, data *[]byte, extension string) *verrors.VError {

	if nil == data {
		return verrors.InvalidRequest(verrors.UserProfilePictureEmptyMessage)
	}

	usr, uerr := dal.GetUser(db, userID)
	if nil != uerr {
		return uerr
	}

	basePath := fmt.Sprintf("data/%s", usr.Database)
	_, err := os.Stat(basePath)
	if nil != err {
		if !os.IsNotExist(err) {
			return verrors.InvalidRequest(err.Error())
		}

		err = os.MkdirAll(basePath, 0775)
		if nil != err {
			return verrors.InvalidRequest(err.Error())
		}
	}

	fileName := fmt.Sprintf("%s/profile_pic%s", basePath, extension)
	ferr := inout.SaveImage(fileName, data)
	if nil != ferr {
		return verrors.Unexpected(ferr.Error())
	}

	uerr = dal.UpdateUserProfilePicture(db, userID, fileName)
	if nil != uerr {
		return uerr
	}

	return nil
}
