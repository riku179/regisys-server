package main

import (
	"context"
	"encoding/base64"
	"github.com/riku179/regisys/app/test"
	"testing"
)

var ctrl, _ = NewJWTController(service)

func TestNewJWTController(t *testing.T) {
	ctx := context.Background()
	basicToken := base64.StdEncoding.EncodeToString([]byte("user:password"))
	basicHeader := "Basic " + basicToken

	//ctx = context.WithValue(ctx, "Authorization")
	resp, _ := test.SigninJWTOK(t, ctx, service, ctrl, true, basicHeader)
	jwtToken := resp.Header().Get("Authorization")
}
