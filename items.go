package main

import (
	"github.com/goadesign/goa"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/models"
)

// ItemsController implements the items resource.
type ItemsController struct {
	*goa.Controller
	DB *models.ItemsDB
}

// NewItemsController creates a items controller.
func NewItemsController(service *goa.Service, db *models.ItemsDB) *ItemsController {
	return &ItemsController{
		Controller: service.NewController("ItemsController"),
		DB:         db,
	}
}

// Add runs the add action.
func (c *ItemsController) Add(ctx *app.AddItemsContext) error {
	// ItemsController_Add: start_implement

	// Put your logic here

	// ItemsController_Add: end_implement
	return nil
}

// Delete runs the delete action.
func (c *ItemsController) Delete(ctx *app.DeleteItemsContext) error {
	// ItemsController_Delete: start_implement

	// Put your logic here

	// ItemsController_Delete: end_implement
	return nil
}

// Modify runs the modify action.
func (c *ItemsController) Modify(ctx *app.ModifyItemsContext) error {
	// ItemsController_Modify: start_implement

	// Put your logic here

	// ItemsController_Modify: end_implement
	return nil
}

// Show runs the show action.
func (c *ItemsController) Show(ctx *app.ShowItemsContext) error {
	// ItemsController_Show: start_implement

	// Put your logic here

	// ItemsController_Show: end_implement
	res := app.GoaExampleRegisysItemsCollection{}
	return ctx.OK(res)
}
