package main

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/riku179/regisys-server/app/test"
)

var jwtCtrl, _ = NewJWTController(service)

func TestSigninJWTOK(t *testing.T) {
	ctx := context.Background()
	basicToken := base64.StdEncoding.EncodeToString([]byte("JohnDoe:foobar"))
	basicHeader := "Basic " + basicToken

	test.SigninJWTOK(t, ctx, service, jwtCtrl, false, basicHeader)
}

func TestSigninJWTUnauthorization(t *testing.T) {
	ctx := context.Background()
	basicToken := base64.StdEncoding.EncodeToString([]byte("Michael:Jackson"))
	basicHeader := "Basic " + basicToken

	test.SigninJWTUnauthorized(t, ctx, service, jwtCtrl, false, basicHeader)
}
