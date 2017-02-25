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

func TestAddUser_Forbidden(t *testing.T) {
	OB := &models.User{Name: "OB", Password: "password", Group: Normal,
		IsMember: false,
	}
	obCtx := user.NewContext(ctx, OB)

	t.Log("非メンバーユーザーが他ユーザーを追加")
	test.AddUserForbidden(
		t, obCtx, service, userCtrl, &app.AddUserPayload{Name: "foo", Password: "bar"})
}

func TestAddUser_NoContent(t *testing.T) {
	//// only MMA member can add non-MMA user
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("メンバーユーザーが他ユーザーを追加")
	test.AddUserNoContent(
		t, normalCtx, service, userCtrl, &app.AddUserPayload{Name: "foo", Password: "bar"})

	// check user exists
	var addUser models.User
	UserDB.Db.Last(&addUser)
	defer UserDB.Delete(ctx, addUser.ID)
	if addUser.Name != "foo" {
		t.Fatalf("Username expected: foo, but received: %+v", addUser.Name)
	}
}

func TestModifyUser_Forbidden(t *testing.T) {
	// Power: Admin > Register > Normal
	// User cannot modify others who have more power than themselves
	// and, User cannot change others to more powerful than themselves
	normal, normalCtx := PrepareUser(Normal)
	register, registerCtx := PrepareUser(Register)
	admin, _ := PrepareUser(Admin)
	defer UserDB.Delete(ctx, register.ID)
	defer UserDB.Delete(ctx, normal.ID)
	defer UserDB.Delete(ctx, admin.ID)

	t.Log("Normal user try other Register -> Admin")
	test.ModifyUserForbidden(
		t, normalCtx, service, userCtrl, register.ID, &app.ModifyUserPayload{
			Group: Register,
		})

	t.Log("Register user try other Admin -> Register")
	test.ModifyUserForbidden(
		t, registerCtx, service, userCtrl, admin.ID, &app.ModifyUserPayload{
			Group: Normal,
		})

	t.Log("Register user try other Normal -> Admin")
	test.ModifyUserForbidden(
		t, registerCtx, service, userCtrl, normal.ID, &app.ModifyUserPayload{
			Group: Admin,
		})
}

func TestModifyUser_NoContent(t *testing.T) {
	// reference comment in TestModifyForbidden()

	normal, _ := PrepareUser(Normal)
	register, registerCtx := PrepareUser(Register)
	admin, adminCtx := PrepareUser(Admin)
	defer UserDB.Delete(ctx, register.ID)
	defer UserDB.Delete(ctx, normal.ID)
	defer UserDB.Delete(ctx, admin.ID)

	t.Log("Register user try other Normal -> Register")
	test.ModifyUserNoContent(
		t, registerCtx, service, userCtrl, normal.ID, &app.ModifyUserPayload{
			Group: Register,
		})
	t.Log("Admin user try other Register -> Admin")
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

func TestModifyUser_NotFound(t *testing.T) {
	admin, adminCtx := PrepareUser(Admin)
	defer UserDB.Delete(ctx, admin.ID)

	t.Log("存在しないユーザーのグループを更新")
	test.ModifyUserNotFound(t, adminCtx, service, userCtrl, 114514, &app.ModifyUserPayload{
		Group: Register,
	})
}

func TestShowUser_OK(t *testing.T) {
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("存在するユーザーを取得")
	test.ShowUserOK(t, normalCtx, service, userCtrl, normal.ID)
}

func TestShowUser_NotFound(t *testing.T) {
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("存在しないユーザーを取得")
	test.ShowUserNotFound(t, normalCtx, service, userCtrl, 114514)
}

func TestShowList_OK(t *testing.T) {
	normal, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("全てのユーザーをの一覧を表示")
	test.ShowListUserOK(t, normalCtx, service, userCtrl)
}

// Prepare User model and context from user's group
func PrepareUser(group string) (usr *models.User, ctx context.Context) {
	usr = &models.User{Name: group, Password: "password", Group: group, IsMember: true}
	UserDB.Add(ctx, usr)
	ctx = user.NewContext(ctx, usr)
	return
}
