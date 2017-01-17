package main

import (
	"github.com/goadesign/goa"
	"github.com/riku179/regisys/app"
)

// GoodsController implements the goods resource.
type GoodsController struct {
	*goa.Controller
}

// NewGoodsController creates a goods controller.
func NewGoodsController(service *goa.Service) *GoodsController {
	return &GoodsController{Controller: service.NewController("GoodsController")}
}

// Add runs the add action.
func (c *GoodsController) Add(ctx *app.AddGoodsContext) error {
	// GoodsController_Add: start_implement

	// Put your logic here

	// GoodsController_Add: end_implement
	return nil
}

// Delete runs the delete action.
func (c *GoodsController) Delete(ctx *app.DeleteGoodsContext) error {
	// GoodsController_Delete: start_implement

	// Put your logic here

	// GoodsController_Delete: end_implement
	return nil
}

// Modify runs the modify action.
func (c *GoodsController) Modify(ctx *app.ModifyGoodsContext) error {
	// GoodsController_Modify: start_implement

	// Put your logic here

	// GoodsController_Modify: end_implement
	return nil
}

// Show runs the show action.
func (c *GoodsController) Show(ctx *app.ShowGoodsContext) error {
	// GoodsController_Show: start_implement

	// Put your logic here

	// GoodsController_Show: end_implement
	res := app.GoaExampleRegisysGoodsCollection{}
	return ctx.OK(res)
}
