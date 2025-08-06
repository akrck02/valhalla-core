package tests

import (
	"testing"

	"github.com/akrck02/valhalla-core/sdk/models"
)

func TestPermissionBits(t *testing.T) {

	var role byte = byte(models.ReadProjectPermission) | byte(models.WriteProjectPermission)
	hasRead := role&byte(models.ReadProjectPermission) == byte(models.ReadProjectPermission)
	hasWrite := role&byte(models.WriteProjectPermission) == byte(models.WriteProjectPermission)
	hasManage := role&byte(models.ManageProjectPermission) == byte(models.ManageProjectPermission)

	if !hasRead || !hasWrite || hasManage {
		t.Fail()
	}
}
