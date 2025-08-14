package schema

import "github.com/akrck02/valhalla-core/sdk/models"

type Project struct {
	Shared      bool   `json:"shared,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	InsertDate  int64  `json:"insert_date,omitempty"`
	UpdateDate  int64  `json:"update_date,omitempty"`
}

type ProjectMember struct {
	UserID      int64                      `json:"user_id,omitempty"`
	ProjectID   int64                      `json:"project_id,omitempty"`
	Permissions []models.ProjectPermission `json:"permissions,omitempty"`
}
