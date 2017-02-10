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
	login_user, _ := user.FromContext(ctx)

	if login_user.IsMember == false {
		return ctx.Forbidden()
	}
	// UserController_Add: start_implement
	hash, _ := bcrypt.GenerateFromPassword([]byte(ctx.Payload.Password), 10)
	res_user := models.User{IsMember: false, Name: ctx.Payload.Name, Password: string(hash)}
	UserDB.Add(ctx, &res_user)
	// UserController_Add: end_implement
	return nil
}

// Modify runs the modify action.
func (c *UserController) Modify(ctx *app.ModifyUserContext) error {
	// login user
	login_user, _ := user.FromContext(ctx)

	// target user requested modify group by login_user
	target_user, err := UserDB.Get(ctx, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}

	// Normal user can do nothing
	if login_user.Group == Normal {
		return ctx.Forbidden()
	}

	// 'Register' user can only upgrade 'Normal' user's group to 'Register'
	if login_user.Group == Register {
		if target_user.Group != Normal || ctx.Payload.Group != Register {
			return ctx.Forbidden()
		}
	}
	// and 'Admin' can do anything

	// UserController_Modify: start_implement
	err = UserDB.UpdateFromModifyUserPayload(ctx, ctx.Payload, ctx.ID)
	if err != nil {
		goa.LogError(ctx, "Failed to access DB", err)
	}
	// UserController_Modify: end_implement
	return nil
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement
	// get requested user from DB
	req_user, err := UserDB.Get(ctx, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	// UserController_Show: end_implement
	res := &app.GoaExampleUser{
		ID:       req_user.ID,
		Group:    req_user.Group,
		Name:     req_user.Name,
		IsMember: req_user.IsMember,
	}
	return ctx.OK(res)
}

// ShowList runs the showList action.
func (c *UserController) ShowList(ctx *app.ShowListUserContext) error {
	//UserController_ShowList: start_implement
	users, err := UserDB.List(ctx)
	if err != nil {
		goa.LogError(ctx, "Failed to access DB", err)
	}
	// UserController_ShowList: end_implement
	//res := app.GoaExampleUserCollection{}
	res := []*app.GoaExampleUser{}
	for _, usr := range users {
		res = append(res, &app.GoaExampleUser{
			ID:       usr.ID,
			Group:    usr.Group,
			Name:     usr.Name,
			IsMember: usr.IsMember,
		})
	}
	return ctx.OK(res)
}
