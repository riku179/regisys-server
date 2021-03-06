package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/models"
	"github.com/riku179/regisys-server/user"
	"golang.org/x/crypto/bcrypt"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
	}
}

// Add runs the add action.
func (c *UserController) Add(ctx *app.AddUserContext) error {
	// Only MMA member can add non member user
	reqUser := user.FromContext(ctx)

	if reqUser.IsMember == false {
		return ctx.Forbidden()
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(ctx.Payload.Password), 10)
	res_user := models.User{IsMember: false, Name: ctx.Payload.Name, Password: string(hash)}
	UserDB.Add(ctx, &res_user)

	return ctx.NoContent()
}

// Modify runs the modify action.
func (c *UserController) Modify(ctx *app.ModifyUserContext) error {
	// login user
	reqUser := user.FromContext(ctx)

	// target user requested modify authority by login_user
	_, err := UserDB.Get(ctx, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}

	// only Register user can do
	if reqUser.IsRegister == false {
		return ctx.Forbidden()
	}

	err = UserDB.UpdateFromModifyUserPayload(ctx, ctx.Payload, ctx.ID)
	if err != nil {
		goa.LogError(ctx, "Failed to access DB", err)
	}
	return ctx.NoContent()
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// get requested user from DB
	reqUser, err := UserDB.Get(ctx, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}

	res := &app.RegisysUser{
		ID:         reqUser.ID,
		IsRegister: reqUser.IsRegister,
		Name:       reqUser.Name,
		IsMember:   reqUser.IsMember,
	}
	return ctx.OK(res)
}

// ShowList runs the showList action.
func (c *UserController) ShowList(ctx *app.ShowListUserContext) error {
	users, err := UserDB.List(ctx)
	if err != nil {
		goa.LogError(ctx, "Failed to access DB", err)
	}

	res := []*app.RegisysUser{}
	for _, usr := range users {
		res = append(res, &app.RegisysUser{
			ID:         usr.ID,
			IsRegister: usr.IsRegister,
			Name:       usr.Name,
			IsMember:   usr.IsMember,
		})
	}
	return ctx.OK(res)
}
