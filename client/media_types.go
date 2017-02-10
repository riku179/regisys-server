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
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// An order (default view)
//
// Identifier: application/vnd.goa.example.orders+json; view=default
type GoaExampleOrders struct {
	// Order date
	Date time.Time `form:"date" json:"date" xml:"date"`
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

// Validate validates the GoaExampleOrders media type instance.
func (mt *GoaExampleOrders) Validate() (err error) {

	if mt.ItemName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_name"))
	}

	return
}

// DecodeGoaExampleOrders decodes the GoaExampleOrders instance encoded in resp body.
func (c *Client) DecodeGoaExampleOrders(resp *http.Response) (*GoaExampleOrders, error) {
	var decoded GoaExampleOrders
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaExampleOrdersCollection is the media type for an array of GoaExampleOrders (default view)
//
// Identifier: application/vnd.goa.example.orders+json; type=collection; view=default
type GoaExampleOrdersCollection []*GoaExampleOrders

// Validate validates the GoaExampleOrdersCollection media type instance.
func (mt GoaExampleOrdersCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeGoaExampleOrdersCollection decodes the GoaExampleOrdersCollection instance encoded in resp body.
func (c *Client) DecodeGoaExampleOrdersCollection(resp *http.Response) (GoaExampleOrdersCollection, error) {
	var decoded GoaExampleOrdersCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// An item (default view)
//
// Identifier: application/vnd.goa.example.regisys.items+json; view=default
type GoaExampleRegisysItems struct {
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
	// Username
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
}

// Validate validates the GoaExampleRegisysItems media type instance.
func (mt *GoaExampleRegisysItems) Validate() (err error) {

	if mt.ItemName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_name"))
	}

	if mt.UserName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user_name"))
	}

	return
}

// DecodeGoaExampleRegisysItems decodes the GoaExampleRegisysItems instance encoded in resp body.
func (c *Client) DecodeGoaExampleRegisysItems(resp *http.Response) (*GoaExampleRegisysItems, error) {
	var decoded GoaExampleRegisysItems
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaExampleRegisysItemsCollection is the media type for an array of GoaExampleRegisysItems (default view)
//
// Identifier: application/vnd.goa.example.regisys.items+json; type=collection; view=default
type GoaExampleRegisysItemsCollection []*GoaExampleRegisysItems

// Validate validates the GoaExampleRegisysItemsCollection media type instance.
func (mt GoaExampleRegisysItemsCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeGoaExampleRegisysItemsCollection decodes the GoaExampleRegisysItemsCollection instance encoded in resp body.
func (c *Client) DecodeGoaExampleRegisysItemsCollection(resp *http.Response) (GoaExampleRegisysItemsCollection, error) {
	var decoded GoaExampleRegisysItemsCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Username and ID (default view)
//
// Identifier: application/vnd.goa.example.token+json; view=default
type GoaExampleToken struct {
	// Group of user
	Group string `form:"group" json:"group" xml:"group"`
	// Unique user ID
	ID int `form:"id" json:"id" xml:"id"`
	// Is member of MMA
	IsMember bool `form:"is_member" json:"is_member" xml:"is_member"`
	// Username
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the GoaExampleToken media type instance.
func (mt *GoaExampleToken) Validate() (err error) {

	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.Group == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "group"))
	}

	return
}

// DecodeGoaExampleToken decodes the GoaExampleToken instance encoded in resp body.
func (c *Client) DecodeGoaExampleToken(resp *http.Response) (*GoaExampleToken, error) {
	var decoded GoaExampleToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Users (default view)
//
// Identifier: application/vnd.goa.example.user+json; view=default
type GoaExampleUser struct {
	// Group of user
	Group string `form:"group" json:"group" xml:"group"`
	// Unique user ID
	ID int `form:"id" json:"id" xml:"id"`
	// Is member of MMA
	IsMember bool `form:"is_member" json:"is_member" xml:"is_member"`
	// Username
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GoaExampleUser media type instance.
func (mt *GoaExampleUser) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Group == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "group"))
	}

	return
}

// DecodeGoaExampleUser decodes the GoaExampleUser instance encoded in resp body.
func (c *Client) DecodeGoaExampleUser(resp *http.Response) (*GoaExampleUser, error) {
	var decoded GoaExampleUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaExampleUserCollection is the media type for an array of GoaExampleUser (default view)
//
// Identifier: application/vnd.goa.example.user+json; type=collection; view=default
type GoaExampleUserCollection []*GoaExampleUser

// Validate validates the GoaExampleUserCollection media type instance.
func (mt GoaExampleUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeGoaExampleUserCollection decodes the GoaExampleUserCollection instance encoded in resp body.
func (c *Client) DecodeGoaExampleUserCollection(resp *http.Response) (GoaExampleUserCollection, error) {
	var decoded GoaExampleUserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
