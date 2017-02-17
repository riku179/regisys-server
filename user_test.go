package main

import (
	"testing"

	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/app/test"
	"github.com/riku179/regisys-server/models"
	"github.com/riku179/regisys-server/user"
)

var userCtrl = NewUserController(service)

func TestAddUserForbidden(t *testing.T) {
	// only MMA member can add OB account
	ctx = user.NewContext(ctx, &models.User{
		IsMember: false,
		Name:     "Mr.OB",
		Password: "foobar",
	})

	test.AddUserForbidden(t, ctx, service, userCtrl, &app.AddUserPayload{Name: "foo", Password: "bar"})
}

func TestAddUserNoContent(t *testing.T) {
	// only MMA member can add non-MMA user
	ctx = user.NewContext(ctx, &models.User{
		IsMember: true,
		Name:     "Mr.MMA",
		Password: "foobar",
	})

	addUserPayload := &app.AddUserPayload{Name: "foo", Password: "bar"}

	test.AddUserNoContent(t, ctx, service, userCtrl, addUserPayload)
}

func TestModifyUserForbidden(t *testing.T) {
	// Power: Admin > Register > Normal
	// User cannot modify others have more power than themselves

	// Normal -> Register
	modifyUserPayload := &app.ModifyUserPayload{Group: ""}
}
