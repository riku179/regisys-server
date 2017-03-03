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
	OB := &models.User{Name: "OB", Password: "password", IsRegister: false,
		IsMember: false,
	}
	obCtx := user.NewContext(ctx, OB)

	t.Log("非メンバーユーザーが他ユーザーを追加")
	test.AddUserForbidden(
		t, obCtx, service, userCtrl, &app.AddUserPayload{Name: "foo", Password: "bar"})
}

func TestAddUser_NoContent(t *testing.T) {
	//// only MMA member can add non-MMA user
	normal, normalCtx := PrepareUser(false)
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
	normal, normalCtx := PrepareUser(false)
	normal2, _ := PrepareUser(false)
	defer UserDB.Delete(ctx, normal.ID)
	defer UserDB.Delete(ctx, normal2.ID)

	t.Log("Normal User can't modify other Normal User to Register User")
	test.ModifyUserForbidden(
		t, normalCtx, service, userCtrl, normal2.ID, &app.ModifyUserPayload{
			IsRegister: true,
		})
}

func TestModifyUser_NoContent(t *testing.T) {
	normal, _ := PrepareUser(false)
	register, registerCtx := PrepareUser(true)
	defer UserDB.Delete(ctx, register.ID)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("Register user can modify Normal User to Register User")
	test.ModifyUserNoContent(
		t, registerCtx, service, userCtrl, normal.ID, &app.ModifyUserPayload{
			IsRegister: true,
		})

	u, err := UserDB.Get(ctx, normal.ID)
	if err != nil {
		t.Fatalf("cought error to get user from DB [Err] %+v", err)
	}
	if u.IsRegister != true {
		t.Fatalf("user's authority didn't change. got: %+v", u.IsRegister)
	}
}

func TestModifyUser_NotFound(t *testing.T) {
	register, registerCtx := PrepareUser(true)
	defer UserDB.Delete(ctx, register.ID)

	t.Log("存在しないユーザーのグループを更新")
	test.ModifyUserNotFound(t, registerCtx, service, userCtrl, 114514, &app.ModifyUserPayload{
		IsRegister: true,
	})
}

func TestShowUser_OK(t *testing.T) {
	normal, normalCtx := PrepareUser(false)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("存在するユーザーを取得")
	test.ShowUserOK(t, normalCtx, service, userCtrl, normal.ID)
}

func TestShowUser_NotFound(t *testing.T) {
	normal, normalCtx := PrepareUser(false)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("存在しないユーザーを取得")
	test.ShowUserNotFound(t, normalCtx, service, userCtrl, 114514)
}

func TestShowList_OK(t *testing.T) {
	normal, normalCtx := PrepareUser(false)
	defer UserDB.Delete(ctx, normal.ID)

	t.Log("全てのユーザーをの一覧を表示")
	test.ShowListUserOK(t, normalCtx, service, userCtrl)
}

// Prepare User model and context from user's authority
func PrepareUser(is_register bool) (usr *models.User, ctx context.Context) {
	if is_register {
		usr = &models.User{Name: "NotRegister", Password: "password", IsRegister: is_register, IsMember: true}
	} else {
		usr = &models.User{Name: "Register", Password: "password", IsRegister: is_register, IsMember: true}
	}
	UserDB.Add(ctx, usr)
	ctx = user.NewContext(ctx, usr)
	return
}
