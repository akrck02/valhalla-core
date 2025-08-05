package dal

import (
	"github.com/akrck02/valhalla-core/sdk/errors"
	"github.com/akrck02/valhalla-core/sdk/models"
)

func CreateProject(project models.Project) (*string, *errors.VError) {
	return nil, errors.TODO()
}

func RemoveProject(id string) *errors.VError {
	return errors.TODO()
}

func UpdateProject(id string, project models.Project) *errors.VError {
	return errors.TODO()
}

func GetProject(id string) (*models.Project, *errors.VError) {
	return nil, errors.TODO()
}

func GetAllProjectsByMember(userId string) ([]models.Project, *errors.VError) {
	return make([]models.Project, 0), errors.TODO()
}
