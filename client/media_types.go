// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/riku179/regisys-server/design
// --out=$(GOPATH)/src/github.com/riku179/regisys-server
// --version=v1.1.0-dirty
//
// API "regisys": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"net/http"

	"github.com/goadesign/goa"
)

// An item (default view)
//
// Identifier: application/vnd.regisys.items+json; view=default
type RegisysItems struct {
	// Unique item ID
	ID int `form:"id" json:"id" xml:"id"`
	// item name
	ItemName string `form:"item_name" json:"item_name" xml:"item_name"`
	// Membership discount
	MemberPrice int `form:"member_price" json:"member_price" xml:"member_price"`
	// item price
	Price int `form:"price" json:"price" xml:"price"`
	// item quantity
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
	// Unique User ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the RegisysItems media type instance.
func (mt *RegisysItems) Validate() (err error) {

	if mt.ItemName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_name"))
	}

	return
}

// DecodeRegisysItems decodes the RegisysItems instance encoded in resp body.
func (c *Client) DecodeRegisysItems(resp *http.Response) (*RegisysItems, error) {
	var decoded RegisysItems
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RegisysItemsCollection is the media type for an array of RegisysItems (default view)
//
// Identifier: application/vnd.regisys.items+json; type=collection; view=default
type RegisysItemsCollection []*RegisysItems

// Validate validates the RegisysItemsCollection media type instance.
func (mt RegisysItemsCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeRegisysItemsCollection decodes the RegisysItemsCollection instance encoded in resp body.
func (c *Client) DecodeRegisysItemsCollection(resp *http.Response) (RegisysItemsCollection, error) {
	var decoded RegisysItemsCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// An order (default view)
//
// Identifier: application/vnd.regisys.orders+json; view=default
type RegisysOrders struct {
	// Order datetime
	Datetime int `form:"datetime" json:"datetime" xml:"datetime"`
	// Unique order ID
	ID int `form:"id" json:"id" xml:"id"`
	// Unique item ID
	ItemID int `form:"item_id" json:"item_id" xml:"item_id"`
	// item name
	ItemName string `form:"item_name" json:"item_name" xml:"item_name"`
	// item price
	Price int `form:"price" json:"price" xml:"price"`
	// item quantity
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
	// Register's user ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the RegisysOrders media type instance.
func (mt *RegisysOrders) Validate() (err error) {

	if mt.ItemName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_name"))
	}

	return
}

// DecodeRegisysOrders decodes the RegisysOrders instance encoded in resp body.
func (c *Client) DecodeRegisysOrders(resp *http.Response) (*RegisysOrders, error) {
	var decoded RegisysOrders
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RegisysOrdersCollection is the media type for an array of RegisysOrders (default view)
//
// Identifier: application/vnd.regisys.orders+json; type=collection; view=default
type RegisysOrdersCollection []*RegisysOrders

// Validate validates the RegisysOrdersCollection media type instance.
func (mt RegisysOrdersCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeRegisysOrdersCollection decodes the RegisysOrdersCollection instance encoded in resp body.
func (c *Client) DecodeRegisysOrdersCollection(resp *http.Response) (RegisysOrdersCollection, error) {
	var decoded RegisysOrdersCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Users (default view)
//
// Identifier: application/vnd.regisys.user+json; view=default
type RegisysUser struct {
	// Unique user ID
	ID int `form:"id" json:"id" xml:"id"`
	// Is member of MMA
	IsMember bool `form:"is_member" json:"is_member" xml:"is_member"`
	// Register or not
	IsRegister bool `form:"is_register" json:"is_register" xml:"is_register"`
	// Username
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the RegisysUser media type instance.
func (mt *RegisysUser) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeRegisysUser decodes the RegisysUser instance encoded in resp body.
func (c *Client) DecodeRegisysUser(resp *http.Response) (*RegisysUser, error) {
	var decoded RegisysUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RegisysUserCollection is the media type for an array of RegisysUser (default view)
//
// Identifier: application/vnd.regisys.user+json; type=collection; view=default
type RegisysUserCollection []*RegisysUser

// Validate validates the RegisysUserCollection media type instance.
func (mt RegisysUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeRegisysUserCollection decodes the RegisysUserCollection instance encoded in resp body.
func (c *Client) DecodeRegisysUserCollection(resp *http.Response) (RegisysUserCollection, error) {
	var decoded RegisysUserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
