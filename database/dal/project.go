package dal

import (
	verrors "github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func CreateProject(project models.Project) (*string, *verrors.VError) {
	return nil, nil
}

func RemoveProject(id string) *verrors.VError {
	return nil
}

func UpdateProject(id string, project models.Project) *verrors.VError {
	return nil
}

func GetProject(id string) (*models.Project, *verrors.VError) {
	return nil, nil
}

func GetAllProjectsByMember(userID string) ([]models.Project, *verrors.VError) {
	return make([]models.Project, 0), nil
}
