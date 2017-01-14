package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"os"
)

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access")
})

var SigninBasicAuth = BasicAuthSecurity("SigninBasicAuth")

var _ = API("regisys", func() {
	Title("Register System for Junk-Ichi in MMA")
	Scheme("http")
	Host(os.Getenv("HOST_ADDR"))
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Metadata("swagger:generate", "false")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/*filepath", "swagger/dist")
})
