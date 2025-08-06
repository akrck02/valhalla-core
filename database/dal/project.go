package dal

import (
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func CreateProject(project models.Project) (*string, *errors.VError) {
	return nil, nil
}

func RemoveProject(id string) *errors.VError {
	return nil
}

func UpdateProject(id string, project models.Project) *errors.VError {
	return nil
}

func GetProject(id string) (*models.Project, *errors.VError) {
	return nil, nil
}

func GetAllProjectsByMember(userId string) ([]models.Project, *errors.VError) {
	return make([]models.Project, 0), nil
}
