package dal

import (
	"github.com/akrck02/valhalla-core/errors"
	"github.com/akrck02/valhalla-core/models"
)

func CreateProject(project models.Project) (*string, error) {
	return nil, errors.TODO()
}

func RemoveProject(id string) error {
	return errors.TODO()
}

func UpdateProject(id string, project models.Project) error {
	return errors.TODO()
}

func GetProject(id string) (*models.Project, error) {
	return nil, errors.TODO()
}

func GetAllProjectsByMember(userId string) ([]models.Project, error) {
	return make([]models.Project, 0), errors.TODO()
}
