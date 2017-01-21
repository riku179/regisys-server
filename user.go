package main

import (
	"github.com/goadesign/goa"
	"github.com/riku179/regisys/app"
	"log"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Add runs the add action.
func (c *UserController) Add(ctx *app.AddUserContext) error {
	// UserController_Add: start_implement

	// Put your logic here

	// UserController_Add: end_implement
	return nil
}

// Modify runs the modify action.
func (c *UserController) Modify(ctx *app.ModifyUserContext) error {
	// UserController_Modify: start_implement

	// Put your logic here

	// UserController_Modify: end_implement
	return nil
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here
	log.Println("user_id:", ctx.Value(ID))

	// UserController_Show: end_implement
	res := &app.GoaExampleUser{}
	return ctx.OK(res)
}

// ShowList runs the showList action.
func (c *UserController) ShowList(ctx *app.ShowListUserContext) error {
	// UserController_ShowList: start_implement

	// Put your logic here

	// UserController_ShowList: end_implement
	res := app.GoaExampleUserCollection{}
	return ctx.OK(res)
}
