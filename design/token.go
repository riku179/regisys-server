package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("jwt", func() {
	Description("This resource uses JWT to secure its endpoints")

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("signin", func() {
		Description("Create a valid JWT")
		Security(SigninBasicAuth)
		Routing(GET("/token"))
		Params(func() {
			Param("is_member", Boolean, "Is member of MMA", func() {
				Example(true)
			})
			Required("is_member")
		})
		Headers(func() {
			Header("Authorization", String, "Basic Auth Header", func() {
				Example("Basic Zm9vOnBhc3N3b3JkCg==")
			})
			Required("Authorization")
		})

		Response(OK, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})
})
