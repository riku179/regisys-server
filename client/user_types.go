// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/riku179/regisys/design
// --out=$(GOPATH)/src/github.com/riku179/regisys
// --version=v1.1.0-dirty
//
// API "regisys": Application User Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import "github.com/goadesign/goa"

// addGoodsPayload user type.
type addGoodsPayload struct {
	// item name
	ItemName *string `form:"item_name,omitempty" json:"item_name,omitempty" xml:"item_name,omitempty"`
	// Membership discount
	MemberPrice *int `form:"member_price,omitempty" json:"member_price,omitempty" xml:"member_price,omitempty"`
	// item price
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// item quantity
	Quantity *int `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
}

// Validate validates the addGoodsPayload type instance.
func (ut *addGoodsPayload) Validate() (err error) {
	if ut.ItemName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_name"))
	}
	if ut.Price == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "price"))
	}
	if ut.MemberPrice == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "member_price"))
	}
	if ut.Quantity == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "quantity"))
	}
	return
}

// Publicize creates AddGoodsPayload from addGoodsPayload
func (ut *addGoodsPayload) Publicize() *AddGoodsPayload {
	var pub AddGoodsPayload
	if ut.ItemName != nil {
		pub.ItemName = *ut.ItemName
	}
	if ut.MemberPrice != nil {
		pub.MemberPrice = *ut.MemberPrice
	}
	if ut.Price != nil {
		pub.Price = *ut.Price
	}
	if ut.Quantity != nil {
		pub.Quantity = *ut.Quantity
	}
	return &pub
}

// AddGoodsPayload user type.
type AddGoodsPayload struct {
	// item name
	ItemName string `form:"item_name" json:"item_name" xml:"item_name"`
	// Membership discount
	MemberPrice int `form:"member_price" json:"member_price" xml:"member_price"`
	// item price
	Price int `form:"price" json:"price" xml:"price"`
	// item quantity
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
}

// Validate validates the AddGoodsPayload type instance.
func (ut *AddGoodsPayload) Validate() (err error) {
	if ut.ItemName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_name"))
	}

	return
}

// addOrderPayload user type.
type addOrderPayload struct {
	// Unique item ID
	ItemID *int `form:"item_id,omitempty" json:"item_id,omitempty" xml:"item_id,omitempty"`
	// item price
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// item quantity
	Quantity *int `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Register's user ID
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the addOrderPayload type instance.
func (ut *addOrderPayload) Validate() (err error) {
	if ut.ItemID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_id"))
	}
	if ut.Quantity == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "quantity"))
	}
	if ut.Price == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "price"))
	}
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user_id"))
	}
	return
}

// Publicize creates AddOrderPayload from addOrderPayload
func (ut *addOrderPayload) Publicize() *AddOrderPayload {
	var pub AddOrderPayload
	if ut.ItemID != nil {
		pub.ItemID = *ut.ItemID
	}
	if ut.Price != nil {
		pub.Price = *ut.Price
	}
	if ut.Quantity != nil {
		pub.Quantity = *ut.Quantity
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// AddOrderPayload user type.
type AddOrderPayload struct {
	// Unique item ID
	ItemID int `form:"item_id" json:"item_id" xml:"item_id"`
	// item price
	Price int `form:"price" json:"price" xml:"price"`
	// item quantity
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
	// Register's user ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the AddOrderPayload type instance.
func (ut *AddOrderPayload) Validate() (err error) {

	return
}

// modifyGoodsPayload user type.
type modifyGoodsPayload struct {
	// Unique item ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// item name
	ItemName *string `form:"item_name,omitempty" json:"item_name,omitempty" xml:"item_name,omitempty"`
	// Membership discount
	MemberPrice *int `form:"member_price,omitempty" json:"member_price,omitempty" xml:"member_price,omitempty"`
	// item price
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// item quantity
	Quantity *int `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
}

// Validate validates the modifyGoodsPayload type instance.
func (ut *modifyGoodsPayload) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// Publicize creates ModifyGoodsPayload from modifyGoodsPayload
func (ut *modifyGoodsPayload) Publicize() *ModifyGoodsPayload {
	var pub ModifyGoodsPayload
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.ItemName != nil {
		pub.ItemName = ut.ItemName
	}
	if ut.MemberPrice != nil {
		pub.MemberPrice = ut.MemberPrice
	}
	if ut.Price != nil {
		pub.Price = ut.Price
	}
	if ut.Quantity != nil {
		pub.Quantity = ut.Quantity
	}
	return &pub
}

// ModifyGoodsPayload user type.
type ModifyGoodsPayload struct {
	// Unique item ID
	ID int `form:"id" json:"id" xml:"id"`
	// item name
	ItemName *string `form:"item_name,omitempty" json:"item_name,omitempty" xml:"item_name,omitempty"`
	// Membership discount
	MemberPrice *int `form:"member_price,omitempty" json:"member_price,omitempty" xml:"member_price,omitempty"`
	// item price
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// item quantity
	Quantity *int `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
}

// modifyUserPayload user type.
type modifyUserPayload struct {
	Group *string `form:"group,omitempty" json:"group,omitempty" xml:"group,omitempty"`
}

// Validate validates the modifyUserPayload type instance.
func (ut *modifyUserPayload) Validate() (err error) {
	if ut.Group == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "group"))
	}
	if ut.Group != nil {
		if !(*ut.Group == "admin" || *ut.Group == "register" || *ut.Group == "nomal") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.group`, *ut.Group, []interface{}{"admin", "register", "nomal"}))
		}
	}
	return
}

// Publicize creates ModifyUserPayload from modifyUserPayload
func (ut *modifyUserPayload) Publicize() *ModifyUserPayload {
	var pub ModifyUserPayload
	if ut.Group != nil {
		pub.Group = *ut.Group
	}
	return &pub
}

// ModifyUserPayload user type.
type ModifyUserPayload struct {
	Group string `form:"group" json:"group" xml:"group"`
}

// Validate validates the ModifyUserPayload type instance.
func (ut *ModifyUserPayload) Validate() (err error) {
	if ut.Group == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "group"))
	}
	if !(ut.Group == "admin" || ut.Group == "register" || ut.Group == "nomal") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.group`, ut.Group, []interface{}{"admin", "register", "nomal"}))
	}
	return
}
