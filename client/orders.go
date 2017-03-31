// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/riku179/regisys-server/design
// --out=$(GOPATH)/src/github.com/riku179/regisys-server
// --version=v1.1.0-dirty
//
// API "regisys": orders Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// AddOrdersPath computes a request path to the add action of orders.
func AddOrdersPath() string {

	return fmt.Sprintf("/orders")
}

// Add order
func (c *Client) AddOrders(ctx context.Context, path string, payload *AddOrderPayload, contentType string) (*http.Response, error) {
	req, err := c.NewAddOrdersRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddOrdersRequest create the request corresponding to the add action endpoint of the orders resource.
func (c *Client) NewAddOrdersRequest(ctx context.Context, path string, payload *AddOrderPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// DeleteOrdersPath computes a request path to the delete action of orders.
func DeleteOrdersPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/orders/%s", param0)
}

// Disable order
func (c *Client) DeleteOrders(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteOrdersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteOrdersRequest create the request corresponding to the delete action endpoint of the orders resource.
func (c *Client) NewDeleteOrdersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowOrdersPath computes a request path to the show action of orders.
func ShowOrdersPath() string {

	return fmt.Sprintf("/orders")
}

// Get orders
func (c *Client) ShowOrders(ctx context.Context, path string, timeEnd int, timeStart int, user *int) (*http.Response, error) {
	req, err := c.NewShowOrdersRequest(ctx, path, timeEnd, timeStart, user)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowOrdersRequest create the request corresponding to the show action endpoint of the orders resource.
func (c *Client) NewShowOrdersRequest(ctx context.Context, path string, timeEnd int, timeStart int, user *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp16 := strconv.Itoa(timeEnd)
	values.Set("time_end", tmp16)
	tmp17 := strconv.Itoa(timeStart)
	values.Set("time_start", tmp17)
	if user != nil {
		tmp18 := strconv.Itoa(*user)
		values.Set("user", tmp18)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
