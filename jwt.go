package main

import (
	"crypto/rsa"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys/app"
	"github.com/riku179/regisys/ldap_auth"
	"github.com/riku179/regisys/models"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

const (
	USERNAME = "username"
	PASSWORD = "password"
	ID       = "id"
)

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead. <- ????
func NewJWTMiddleware() (goa.Middleware, error) {
	keys, err := loadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(jwt.NewSimpleResolver(keys), checkUser(), app.NewJWTSecurity()), nil
}

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
	UserDB     *models.UserDB
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service, UserDB *models.UserDB) (*JWTController, error) {
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
		UserDB:     UserDB,
	}, nil
}

func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}

			ctx = context.WithValue(ctx, USERNAME, user)
			ctx = context.WithValue(ctx, PASSWORD, pass)

			// Proceed
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}

// Signin runs the signin action.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// JWTController_Signin: start_implement
	username := ctx.Value(USERNAME).(string)
	var user models.User

	if ctx.IsMember {
		//	Authenticate with LDAP
		if ldap_auth.LdapAuth() != nil {
			return ErrUnauthorized("Unknown user")
		}
		err := c.UserDB.Db.Where("name = ?", username).First(&user).Error
		if err == gorm.ErrRecordNotFound {
			user = models.User{Group: "none", IsMember: true, Name: username}
			c.UserDB.Add(ctx, &user)
		}
	} else {
		err := c.UserDB.Db.Where("name = ?", username).First(&user).Error
		if err == gorm.ErrRecordNotFound {
			return ErrUnauthorized("Unknown user")
		}
	}

	// Generate JWT
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in60m := time.Now().Add(time.Duration(60) * time.Minute).Unix()
	token.Claims = jwtgo.MapClaims{
		"exp":    in60m,             // time when the token will expire (60 minutes from now)
		"iat":    time.Now().Unix(), // when the token was issued/created (now)
		"nbf":    2,                 // time before which the token is not yet valid (2 minutes ago)
		"sub":    user.ID,           // the subject/principal is whom the token is about
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
		ID:       user.ID,
		Username: user.Name,
	}
	return ctx.OK(res)
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func loadJWTPublicKeys() ([]jwt.Key, error) {
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

func checkUser() goa.Middleware {
	validate := func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve the token claims
			token := jwt.ContextJWT(ctx)
			claims := token.Claims.(jwtgo.MapClaims)

			// Use the claims to authorize
			subject := claims["sub"]
			if subject == nil {
				return errValidationFailed("Failed to Validate")
			}
			ctx = context.WithValue(ctx, ID, subject)
			return h(ctx, rw, req)
		}
	}
	fm, _ := goa.NewMiddleware(validate)
	return fm
}
