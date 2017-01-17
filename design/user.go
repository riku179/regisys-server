package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/user")

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("add", func() {
		Description("Add user for NOT MMA member)")
		Security(SigninBasicAuth) // username,pwを受け付けるため。JWTと共存できないならダメ
		Routing(POST(""))

		Response(NoContent)
		Response(Forbidden)
	})

	Action("modify", func() {
		Description("Modify group of user")
		Routing(PUT("/:id"))
		Payload(ModifyUserPayload)

		Response(NoContent)
		Response(NotFound)
		Response(Forbidden)
	})

	Action("show", func() {
		Description("Show one user")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "Unique user ID")
		})

		Response(OK, func() {
			Media(ShowUserMedia)
		})
		Response(NotFound)
	})

	Action("showList", func() {
		Description("Show users list")
		Routing(GET("/list"))

		Response(OK, func() {
			Media(CollectionOf(ShowUserMedia))
		})
	})
})

var ShowUserMedia = MediaType("application/vnd.goa.example.user+json", func() {
	Description("Users")
	Attributes(func() {
		Attribute("id", Integer, "Unique user ID")
		Attribute("name", String, "Username")
		Attribute("group", String, "Group of user")
		Required("id", "name", "group")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("group")
	})
})

var ModifyUserPayload = Type("ModifyUserPayload", func() {
	Member("group", String, func() {
		Enum("admin", "register", "nomal")
	})
	Required("group")
})
