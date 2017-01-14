package main

import (
	"crypto/rsa"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/riku179/regisys/app"
	"github.com/riku179/regisys/ldap_auth"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

const (
	username = "username"
)

func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}
			// LDAP Auth
			if ldap_auth.LdapAuth() != nil {
				goa.LogInfo(ctx, "invalid username or password")
				return ErrUnauthorized("invalid username or password")
			}
			ctx = context.WithValue(ctx, username, user)

			// Proceed
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead. <- ????
func NewJWTMiddleware() (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(jwt.NewSimpleResolver(keys), nil, app.NewJWTSecurity()), nil
}

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service) (*JWTController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}

	return &JWTController{
		Controller: service.NewController("JWTController"),
		privateKey: privKey,
	}, nil
}

// Signin runs the signin action.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// JWTController_Signin: start_implement
	user := ctx.Value(username).(string)

	// Generate JWT
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	token.Claims = jwtgo.MapClaims{
		"exp":    in10m,             // time when the token will expire (10 minutes from now)
		"iat":    time.Now().Unix(), // when the token was issued/created (now)
		"nbf":    2,                 // time before which the token is not yet valid (2 minutes ago)
		"sub":    user,              // the subject/principal is whom the token is about
		"scopes": "api:access",      // token scope - not a standard claim
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// JWTController_Signin: end_implement
	res := &app.GoaExampleToken{
		ID:       11111,
		Username: user,
	}
	return ctx.OK(res)
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() ([]jwt.Key, error) {
	keyFiles, err := filepath.Glob("./jwtkey/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]jwt.Key, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}
