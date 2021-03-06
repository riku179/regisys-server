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

// User Relational Model
type User struct {
	ID               int     `gorm:"primary_key"` // Unique user ID
	IsMember         bool    `sql:"not null;default:false"`
	IsRegister       bool    `sql:"not null;default:false"`
	Items            []Items // has many Items
	Name             string  `sql:"not null;default:false"`
	Password         string
	RegisteredOrders []Orders   // has many Orders
	CreatedAt        time.Time  // timestamp
	DeletedAt        *time.Time // timestamp
	UpdatedAt        time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db *gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id int) (*User, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error

	ListRegisysUser(ctx context.Context) []*app.RegisysUser
	OneRegisysUser(ctx context.Context, id int) (*app.RegisysUser, error)

	UpdateFromAddUserPayload(ctx context.Context, payload *app.AddUserPayload, id int) error

	UpdateFromModifyUserPayload(ctx context.Context, payload *app.ModifyUserPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// CRUD Functions

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx context.Context, id int) (*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "get"}, time.Now())

	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of User
func (m *UserDB) List(ctx context.Context) ([]*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "list"}, time.Now())

	var objs []*User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserDB) Add(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding User", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserDB) Update(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating User", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "delete"}, time.Now())

	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting User", "error", err.Error())
		return err
	}

	return nil
}

// UserFromAddUserPayload Converts source AddUserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromAddUserPayload(payload *app.AddUserPayload) *User {
	user := &User{}
	user.Name = payload.Name
	user.Password = payload.Password

	return user
}

// UpdateFromAddUserPayload applies non-nil changes from AddUserPayload to the model and saves it
func (m *UserDB) UpdateFromAddUserPayload(ctx context.Context, payload *app.AddUserPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromaddUserPayload"}, time.Now())

	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}
	obj.Name = payload.Name
	obj.Password = payload.Password

	err = m.Db.Save(&obj).Error
	return err
}

// UserFromModifyUserPayload Converts source ModifyUserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromModifyUserPayload(payload *app.ModifyUserPayload) *User {
	user := &User{}
	user.IsRegister = payload.IsRegister

	return user
}

// UpdateFromModifyUserPayload applies non-nil changes from ModifyUserPayload to the model and saves it
func (m *UserDB) UpdateFromModifyUserPayload(ctx context.Context, payload *app.ModifyUserPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefrommodifyUserPayload"}, time.Now())

	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}
	obj.IsRegister = payload.IsRegister

	err = m.Db.Save(&obj).Error
	return err
}
