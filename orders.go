package main

import (
	"github.com/goadesign/goa"
	"github.com/riku179/regisys/app"
	"github.com/riku179/regisys/models"
)

// OrdersController implements the orders resource.
type OrdersController struct {
	*goa.Controller
	DB *models.OrdersDB
}

// NewOrdersController creates a orders controller.
func NewOrdersController(service *goa.Service, db *models.OrdersDB) *OrdersController {
	return &OrdersController{
		Controller: service.NewController("OrdersController"),
		DB:         db,
	}
}

// Add runs the add action.
func (c *OrdersController) Add(ctx *app.AddOrdersContext) error {
	// OrdersController_Add: start_implement

	// Put your logic here

	// OrdersController_Add: end_implement
	return nil
}

// Delete runs the delete action.
func (c *OrdersController) Delete(ctx *app.DeleteOrdersContext) error {
	// OrdersController_Delete: start_implement

	// Put your logic here

	// OrdersController_Delete: end_implement
	return nil
}

// Show runs the show action.
func (c *OrdersController) Show(ctx *app.ShowOrdersContext) error {
	// OrdersController_Show: start_implement

	// Put your logic here

	// OrdersController_Show: end_implement
	res := app.GoaExampleOrdersCollection{}
	return ctx.OK(res)
}
