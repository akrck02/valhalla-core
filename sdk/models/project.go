package models

type Project struct {
	Name        string `json:"nm,omitempty"`
	Description string `json:"ds,omitempty"`
	InsertDate  int64  `json:"ind,omitempty"`
	UpdateDate  int64  `json:"ud,omitempty"`
}

type ProjectMember struct {
	UserId      int64               `json:"uid,omitempty"`
	ProjectId   int64               `json:"pid,omitempty"`
	Permissions []ProjectPermission `json:"pms,omitempty"`
}

type ProjectPermission byte

const (
	ReadProjectPermission ProjectPermission = 1 << iota
	WriteProjectPermission
	ManageProjectPermission
)
