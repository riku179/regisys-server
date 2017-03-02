package main

import (
	"encoding/base64"
	"testing"

	"github.com/riku179/regisys-server/app/test"
	"github.com/riku179/regisys-server/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtCtrl, _ = NewJWTController(service)

func TestSigninOKbyMember(t *testing.T) {
	normal, _ := PrepareUser(Normal)
	defer UserDB.Delete(ctx, normal.ID)

	basicHeader := genBasicHeader(normal.Name, normal.Password)

	test.SigninJWTOK(t, ctx, service, jwtCtrl, true, basicHeader)
}

func TestSigninOKbyNotMember(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	OB := &models.User{Name: "OB", Password: string(hash), Group: Normal,
		IsMember: false}
	UserDB.Add(ctx, OB)
	defer UserDB.Delete(ctx, OB.ID)

	basicHeader := genBasicHeader(OB.Name, "password")

	test.SigninJWTOK(t, ctx, service, jwtCtrl, false, basicHeader)
}

func TestSigninUnauthorization(t *testing.T) {
	basicHeader := genBasicHeader("UnauthorizedUser", "hogefuga")

	test.SigninJWTUnauthorized(t, ctx, service, jwtCtrl, false, basicHeader)
}

func genBasicHeader(username, password string) (header string) {
	header = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	return
}
