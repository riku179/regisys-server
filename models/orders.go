// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/riku179/regisys-server/design
// --out=$(GOPATH)/src/github.com/riku179/regisys-server
// --version=v1.1.0-dirty
//
// API "regisys": Models
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"golang.org/x/net/context"
)

// Orders Relational Model
type Orders struct {
	ID            int `gorm:"primary_key"` // Unique order ID
	IsMemberPrice bool
	ItemID        int `sql:"not null;default:false"` // has many Orders
	Price         int `sql:"not null;default:false"`
	Quantity      int `sql:"not null;default:false"`
	UpdatedAt     time.Time
	UserID        int        // has many Orders
	CreatedAt     time.Time  // timestamp
	DeletedAt     *time.Time // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Orders) TableName() string {
	return "orders"

}

// OrdersDB is the implementation of the storage interface for
// Orders.
type OrdersDB struct {
	Db *gorm.DB
}

// NewOrdersDB creates a new storage type.
func NewOrdersDB(db *gorm.DB) *OrdersDB {
	return &OrdersDB{Db: db}
}

// DB returns the underlying database.
func (m *OrdersDB) DB() interface{} {
	return m.Db
}

// OrdersStorage represents the storage interface.
type OrdersStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Orders, error)
	Get(ctx context.Context, id int) (*Orders, error)
	Add(ctx context.Context, orders *Orders) error
	Update(ctx context.Context, orders *Orders) error
	Delete(ctx context.Context, id int) error

	ListRegisysOrders(ctx context.Context) []*app.RegisysOrders
	OneRegisysOrders(ctx context.Context, id int) (*app.RegisysOrders, error)

	UpdateFromAddOrderPayload(ctx context.Context, payload *app.AddOrderPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *OrdersDB) TableName() string {
	return "orders"

}

// CRUD Functions

// Get returns a single Orders as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *OrdersDB) Get(ctx context.Context, id int) (*Orders, error) {
	defer goa.MeasureSince([]string{"goa", "db", "orders", "get"}, time.Now())

	var native Orders
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Orders
func (m *OrdersDB) List(ctx context.Context) ([]*Orders, error) {
	defer goa.MeasureSince([]string{"goa", "db", "orders", "list"}, time.Now())

	var objs []*Orders
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *OrdersDB) Add(ctx context.Context, model *Orders) error {
	defer goa.MeasureSince([]string{"goa", "db", "orders", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Orders", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *OrdersDB) Update(ctx context.Context, model *Orders) error {
	defer goa.MeasureSince([]string{"goa", "db", "orders", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Orders", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *OrdersDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "orders", "delete"}, time.Now())

	var obj Orders

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Orders", "error", err.Error())
		return err
	}

	return nil
}

// OrdersFromAddOrderPayload Converts source AddOrderPayload to target Orders model
// only copying the non-nil fields from the source.
func OrdersFromAddOrderPayload(payload *app.AddOrderPayload) *Orders {
	orders := &Orders{}
	orders.IsMemberPrice = payload.IsMemberPrice
	orders.ItemID = payload.ItemID
	orders.Quantity = payload.Quantity

	return orders
}

// UpdateFromAddOrderPayload applies non-nil changes from AddOrderPayload to the model and saves it
func (m *OrdersDB) UpdateFromAddOrderPayload(ctx context.Context, payload *app.AddOrderPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "orders", "updatefromaddOrderPayload"}, time.Now())

	var obj Orders
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Orders", "error", err.Error())
		return err
	}
	obj.IsMemberPrice = payload.IsMemberPrice
	obj.ItemID = payload.ItemID
	obj.Quantity = payload.Quantity

	err = m.Db.Save(&obj).Error
	return err
}
