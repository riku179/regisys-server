package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/ldap_auth"
	"github.com/riku179/regisys-server/models"
	"github.com/riku179/regisys-server/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
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

func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			return h(ctx, rw, req)
		}
	}
}

// Signin runs the signin action.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// JWTController_Signin: start_implement
	username, password, ok := ctx.BasicAuth()
	if !ok {
		goa.LogInfo(ctx, "failed basic auth")
		return ctx.Unauthorized()
	} else if username == "" {
		goa.LogInfo(ctx, "missing username in basicauth header")
		return ctx.Unauthorized()
	}

	// User data(id,name,is_member) is bound to this User instance
	var reqUser models.User

	if ctx.IsMember {
		//	Authenticate with LDAP
		if ldap_auth.LdapAuth() != nil {
			return ctx.Unauthorized()
		}
		err := UserDB.Db.Where("name = ?", username).First(&reqUser).Error
		if err == gorm.ErrRecordNotFound {
			reqUser = models.User{IsMember: true, Name: username}
			UserDB.Add(ctx, &reqUser)
		}
	} else {
		// Authenticate with username and password
		err := UserDB.Db.Where("name = ?", username).First(&reqUser).Error
		if err == gorm.ErrRecordNotFound {
			return ctx.Unauthorized()
		}
		err = bcrypt.CompareHashAndPassword([]byte(reqUser.Password), []byte(password))
		if err != nil {
			return ctx.Unauthorized()
		}
	}

	// Generate JWT
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	// token expire 24h after
	in60m := time.Now().Add(time.Duration(1440) * time.Minute).Unix()
	token.Claims = jwtgo.MapClaims{
		"exp":         in60m,              // time when the token will expire (60 minutes from now)
		"iat":         time.Now().Unix(),  // when the token was issued/created (now)
		"sub":         reqUser.ID,         // the subject/principal is whom the token is about
		"scopes":      "api:access",       // token scope - not a standard claim
		"is_register": reqUser.IsRegister, // is register of user - not a standard claim
		"user_name":   reqUser.Name,       // username - not a standard claim
		"is_member":   reqUser.IsMember,   // is member of MMA - not a standard claim
	}
	// Sign token by private key
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// JWTController_Signin: end_implement
	return nil
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
		return nil, fmt.Errorf("couldn't load public keys for JWT security%s", "")
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
			ctx = user.NewContext(ctx, &models.User{
				ID:         int(subject.(float64)),
				Name:       claims["user_name"].(string),
				IsRegister: claims["is_register"].(bool),
				IsMember:   claims["is_member"].(bool),
			})
			return h(ctx, rw, req)
		}
	}
	fm, _ := goa.NewMiddleware(validate)
	return fm
}
