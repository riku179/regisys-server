// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/riku179/regisys-server/design
// --out=$(GOPATH)/src/github.com/riku179/regisys-server
// --version=v1.1.0-dirty
//
// API "regisys": Model Helpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListGoaExampleOrders returns an array of view: default.
func (m *OrdersDB) ListGoaExampleOrders(ctx context.Context) []*app.GoaExampleOrders {
	defer goa.MeasureSince([]string{"goa", "db", "goaExampleOrders", "listgoaExampleOrders"}, time.Now())

	var native []*Orders
	var objs []*app.GoaExampleOrders
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Orders", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.OrdersToGoaExampleOrders())
	}

	return objs
}

// OrdersToGoaExampleOrders loads a Orders and builds the default view of media type GoaExampleOrders.
func (m *Orders) OrdersToGoaExampleOrders() *app.GoaExampleOrders {
	orders := &app.GoaExampleOrders{}
	orders.ID = m.ID
	orders.ItemID = m.ItemID
	orders.Price = m.Price
	orders.Quantity = m.Quantity
	orders.UserID = m.UserID

	return orders
}

// OneGoaExampleOrders loads a Orders and builds the default view of media type GoaExampleOrders.
func (m *OrdersDB) OneGoaExampleOrders(ctx context.Context, id int) (*app.GoaExampleOrders, error) {
	defer goa.MeasureSince([]string{"goa", "db", "goaExampleOrders", "onegoaExampleOrders"}, time.Now())

	var native Orders
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Orders", "error", err.Error())
		return nil, err
	}

	view := *native.OrdersToGoaExampleOrders()
	return &view, err
}
