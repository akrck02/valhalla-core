package services

import (
	"testing"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/mock"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
)

func TestCreateProject(t *testing.T) {

	client := db.CreateClient()
	conn := db.Connect(*client)
	defer db.Disconnect(*client, conn)

	user := RegisterMockTestUser(t, conn, client)
	CreateMockTestProjectWithUser(t, conn, client, user)
	DeleteTestUser(t, conn, client, user)
}

func TestCreateProjectWithoutOwner(t *testing.T) {

	client := db.CreateClient()
	conn := db.Connect(*client)
	defer db.Disconnect(*client, conn)

	project := models.Project{
		Name:        "Test Project",
		Description: "Test Description",
	}

	CreateTestProjectWithError(t, conn, client, &project, utils.HTTP_STATUS_BAD_REQUEST, error.EMPTY_PROJECT_OWNER)
}

func TestCreateProjectWithoutName(t *testing.T) {

	client := db.CreateClient()
	conn := db.Connect(*client)
	defer db.Disconnect(*client, conn)

	project := &models.Project{
		Description: mock.ProjectDescription(),
		Owner:       mock.Email(),
	}

	CreateTestProjectWithError(t, conn, client, project, utils.HTTP_STATUS_BAD_REQUEST, error.EMPTY_PROJECT_NAME)
}
func TestCreateProjectWithoutDescription(t *testing.T) {

	client := db.CreateClient()
	conn := db.Connect(*client)
	defer db.Disconnect(*client, conn)

	project := &models.Project{
		Name:  mock.ProjectName(),
		Owner: mock.Email(),
	}

	CreateTestProjectWithError(t, conn, client, project, utils.HTTP_STATUS_BAD_REQUEST, error.EMPTY_PROJECT_DESCRIPTION)
}

func TestCreateProjectThatAlreadyExists(t *testing.T) {

	client := db.CreateClient()
	conn := db.Connect(*client)
	defer db.Disconnect(*client, conn)

	user := RegisterMockTestUser(t, conn, client)
	project := CreateMockTestProjectWithUser(t, conn, client, user)

	CreateTestProjectWithError(t, conn, client, project, utils.HTTP_STATUS_CONFLICT, error.PROJECT_ALREADY_EXISTS)
	DeleteTestUser(t, conn, client, user)
}

func TestGetUserProjects(t *testing.T) {

	client := db.CreateClient()
	conn := db.Connect(*client)
	defer db.Disconnect(*client, conn)

	user := RegisterMockTestUser(t, conn, client)
	project := CreateMockTestProjectWithUser(t, conn, client, user)
	project2 := CreateMockTestProjectWithUser(t, conn, client, user)

	projects := GetUserProjects(conn, client, user.Email)

	if len(projects) == 0 {
		t.Errorf("No projects found for user: %v", user.Email)
	}

	if len(projects) != 2 {
		t.Errorf("Incorrect number of projects found for user: %v", user.Email)
	}

	if projects[0].Name != project.Name {
		t.Errorf("Incorrect project found: %v", projects[0].Name)
	}

	if projects[1].Name != project2.Name {
		t.Errorf("Incorrect project found: %v", projects[1].Name)
	}

	DeleteTestUser(t, conn, client, user)
}
