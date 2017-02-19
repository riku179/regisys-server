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

// Items Relational Model
type Items struct {
	ID          int        `gorm:"primary_key"` // Unique item ID
	ItemName    string     `sql:"not null;default:false"`
	MemberPrice int        `sql:"not null;default:false"`
	Orders      []Orders   // has many Orders
	Price       int        `sql:"not null;default:false"`
	Quantity    int        `sql:"not null;default:false"`
	UserID      int        // has many Items
	CreatedAt   time.Time  // timestamp
	DeletedAt   *time.Time // timestamp
	UpdatedAt   time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Items) TableName() string {
	return "items"

}

// ItemsDB is the implementation of the storage interface for
// Items.
type ItemsDB struct {
	Db *gorm.DB
}

// NewItemsDB creates a new storage type.
func NewItemsDB(db *gorm.DB) *ItemsDB {
	return &ItemsDB{Db: db}
}

// DB returns the underlying database.
func (m *ItemsDB) DB() interface{} {
	return m.Db
}

// ItemsStorage represents the storage interface.
type ItemsStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Items, error)
	Get(ctx context.Context, id int) (*Items, error)
	Add(ctx context.Context, items *Items) error
	Update(ctx context.Context, items *Items) error
	Delete(ctx context.Context, id int) error

	ListRegisysItems(ctx context.Context) []*app.RegisysItems
	OneRegisysItems(ctx context.Context, id int) (*app.RegisysItems, error)

	UpdateFromAddGoodsPayload(ctx context.Context, payload *app.AddGoodsPayload, id int) error

	UpdateFromModifyGoodsPayload(ctx context.Context, payload *app.ModifyGoodsPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ItemsDB) TableName() string {
	return "items"

}

// CRUD Functions

// Get returns a single Items as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ItemsDB) Get(ctx context.Context, id int) (*Items, error) {
	defer goa.MeasureSince([]string{"goa", "db", "items", "get"}, time.Now())

	var native Items
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Items
func (m *ItemsDB) List(ctx context.Context) ([]*Items, error) {
	defer goa.MeasureSince([]string{"goa", "db", "items", "list"}, time.Now())

	var objs []*Items
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *ItemsDB) Add(ctx context.Context, model *Items) error {
	defer goa.MeasureSince([]string{"goa", "db", "items", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Items", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *ItemsDB) Update(ctx context.Context, model *Items) error {
	defer goa.MeasureSince([]string{"goa", "db", "items", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Items", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ItemsDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "items", "delete"}, time.Now())

	var obj Items

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Items", "error", err.Error())
		return err
	}

	return nil
}

// ItemsFromAddGoodsPayload Converts source AddGoodsPayload to target Items model
// only copying the non-nil fields from the source.
func ItemsFromAddGoodsPayload(payload *app.AddGoodsPayload) *Items {
	items := &Items{}
	items.ItemName = payload.ItemName
	items.MemberPrice = payload.MemberPrice
	items.Price = payload.Price
	items.Quantity = payload.Quantity

	return items
}

// UpdateFromAddGoodsPayload applies non-nil changes from AddGoodsPayload to the model and saves it
func (m *ItemsDB) UpdateFromAddGoodsPayload(ctx context.Context, payload *app.AddGoodsPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "items", "updatefromaddGoodsPayload"}, time.Now())

	var obj Items
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Items", "error", err.Error())
		return err
	}
	obj.ItemName = payload.ItemName
	obj.MemberPrice = payload.MemberPrice
	obj.Price = payload.Price
	obj.Quantity = payload.Quantity

	err = m.Db.Save(&obj).Error
	return err
}

// ItemsFromModifyGoodsPayload Converts source ModifyGoodsPayload to target Items model
// only copying the non-nil fields from the source.
func ItemsFromModifyGoodsPayload(payload *app.ModifyGoodsPayload) *Items {
	items := &Items{}
	items.ID = payload.ID
	if payload.ItemName != nil {
		items.ItemName = *payload.ItemName
	}
	if payload.MemberPrice != nil {
		items.MemberPrice = *payload.MemberPrice
	}
	if payload.Price != nil {
		items.Price = *payload.Price
	}
	if payload.Quantity != nil {
		items.Quantity = *payload.Quantity
	}

	return items
}

// UpdateFromModifyGoodsPayload applies non-nil changes from ModifyGoodsPayload to the model and saves it
func (m *ItemsDB) UpdateFromModifyGoodsPayload(ctx context.Context, payload *app.ModifyGoodsPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "items", "updatefrommodifyGoodsPayload"}, time.Now())

	var obj Items
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Items", "error", err.Error())
		return err
	}
	obj.ID = payload.ID
	if payload.ItemName != nil {
		obj.ItemName = *payload.ItemName
	}
	if payload.MemberPrice != nil {
		obj.MemberPrice = *payload.MemberPrice
	}
	if payload.Price != nil {
		obj.Price = *payload.Price
	}
	if payload.Quantity != nil {
		obj.Quantity = *payload.Quantity
	}

	err = m.Db.Save(&obj).Error
	return err
}
