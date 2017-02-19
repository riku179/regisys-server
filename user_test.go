package main

import (
	"context"
	"testing"

	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/app/test"
	"github.com/riku179/regisys-server/models"
	"github.com/riku179/regisys-server/user"
)

var userCtrl = NewUserController(service)

func TestAddForbidden(t *testing.T) {
	// only MMA member can add OB account
	OB := &models.User{Name: "OB", Password: "password", Group: Normal,
		IsMember: false,
	}
	obCtx := user.NewContext(ctx, OB)

	test.AddUserForbidden(
		t, obCtx, service, userCtrl, &app.AddUserPayload{Name: "foo", Password: "bar"})
}

func TestAddNoContent(t *testing.T) {
	//// only MMA member can add non-MMA user
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	test.AddUserNoContent(
		t, normalCtx, service, userCtrl, &app.AddUserPayload{Name: "foo", Password: "bar"})

	// check user exists
	fooID, err := getUserID("foo")
	defer UserDB.Delete(ctx, fooID)
	if err != nil {
		t.Fatalf("user not found in DB [Err] %+v", err)
	}
}

func TestModifyForbidden(t *testing.T) {
	// Power: Admin > Register > Normal
	// User cannot modify others who have more power than themselves
	// and, User cannot change others to more powerful than themselves
	normal, normalCtx := PrepareUser(Normal)
	register, registerCtx := PrepareUser(Register)
	admin, _ := PrepareUser(Admin)
	defer UserDB.Delete(ctx, register.ID)
	defer UserDB.Delete(ctx, normal.ID)
	defer UserDB.Delete(ctx, admin.ID)
	// Normal user try other Register -> Admin
	test.ModifyUserForbidden(
		t, normalCtx, service, userCtrl, register.ID, &app.ModifyUserPayload{
			Group: Register,
		})

	// Register user try other Admin -> Register
	test.ModifyUserForbidden(
		t, registerCtx, service, userCtrl, admin.ID, &app.ModifyUserPayload{
			Group: Normal,
		})

	// Register user try other Normal -> Admin
	test.ModifyUserForbidden(
		t, registerCtx, service, userCtrl, normal.ID, &app.ModifyUserPayload{
			Group: Admin,
		})
}

func TestModifyNoContent(t *testing.T) {
	// reference comment in TestModifyForbidden()

	normal, _ := PrepareUser(Normal)
	register, registerCtx := PrepareUser(Register)
	admin, adminCtx := PrepareUser(Admin)
	defer UserDB.Delete(ctx, register.ID)
	defer UserDB.Delete(ctx, normal.ID)
	defer UserDB.Delete(ctx, admin.ID)

	// Register user try other Normal -> Register
	test.ModifyUserNoContent(
		t, registerCtx, service, userCtrl, normal.ID, &app.ModifyUserPayload{
			Group: Register,
		})
	//Admin user try other Register -> Admin
	test.ModifyUserNoContent(
		t, adminCtx, service, userCtrl, register.ID, &app.ModifyUserPayload{
			Group: Admin,
		})

	u, err := UserDB.Get(ctx, normal.ID)
	if err != nil {
		t.Fatalf("cought error to get user from DB [Err] %+v", err)
	}
	if u.Group != Register {
		t.Fatalf("user's group didn't change. got: %+v", u.Group)
	}

	u, err = UserDB.Get(ctx, register.ID)
	if err != nil {
		t.Fatalf("cought error to get user from DB [Err] %+v", err)
	}
	if u.Group != Admin {
		t.Fatalf("user's group didn't change. got: %+v", u.Group)
	}
}

func TestModifyNotFound(t *testing.T) {
	admin, adminCtx := PrepareUser(Admin)
	defer UserDB.Delete(ctx, admin.ID)

	// user of ID:114514 doesn't exist
	test.ModifyUserNotFound(t, adminCtx, service, userCtrl, 114514, &app.ModifyUserPayload{
		Group: Register,
	})
}

func TestShowUserOK(t *testing.T) {
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	test.ShowUserOK(t, normalCtx, service, userCtrl, normal.ID)
}

func TestShowNotFound(t *testing.T) {
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	// user of ID:114514 doesn't exist
	test.ShowUserNotFound(t, normalCtx, service, userCtrl, 114514)
}

func TestShowListOK(t *testing.T) {
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	test.ShowListUserOK(t, normalCtx, service, userCtrl)
}

func PrepareUser(group string) (usr *models.User, ctx context.Context) {
	usr = &models.User{Name: group, Password: "password", Group: group, IsMember: true}
	UserDB.Add(ctx, usr)
	ctx = user.NewContext(ctx, usr)
	return
}

func getUserID(username string) (int, error) {
	var register = models.User{}
	err := UserDB.Db.Where("name = ?", username).First(&register).Error
	if err != nil {
		return 0, err
	}
	return register.ID, nil
}
