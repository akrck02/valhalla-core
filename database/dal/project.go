package dal

import (
	"database/sql"

	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func CreateProject(db *sql.DB, userID int64, project models.Project) (*string, *verrors.VError) {

	return nil, nil
}

func RemoveProject(db *sql.DB, projectID int64) *verrors.VError {
	return nil
}

func UpdateProject(db *sql.DB, projectID int64, project models.Project) *verrors.VError {
	return nil
}

func GetProject(db *sql.DB, projectID int64) (*models.Project, *verrors.VError) {
	return nil, nil
}

func GetAllProjectsByMember(db *sql.DB, userID int64) ([]models.Project, *verrors.VError) {
	return make([]models.Project, 0), nil
}
