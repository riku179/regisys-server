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
		Routing(POST(""))
		Payload(AddUserPayload)

		Response(NoContent)
		Response(Forbidden)
	})

	Action("modify", func() {
		Description("Modify is_register of user")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id", Integer, "Unique user ID", func() {
				Example(1001)
			})
		})
		Payload(ModifyUserPayload)
		Response(NoContent)
		Response(NotFound)
		Response(Forbidden)
	})

	Action("show", func() {
		Description("Show one user")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "Unique user ID", func() {
				Example(1001)
			})
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

var ShowUserMedia = MediaType("application/vnd.regisys.user+json", func() {
	Description("Users")
	Attributes(func() {
		Attribute("id", Integer, "Unique user ID", func() {
			Example(1001)
		})
		Attribute("name", String, "Username", func() {
			Example("Linus Benedict Torvalds")
		})

		Attribute("is_register", Boolean, "Register or not", func() {
			Example(false)
		})
		Attribute("is_member", Boolean, "Is member of MMA", func() {
			Example(true)
		})
		Required("id", "name", "is_register", "is_member")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("is_register")
		Attribute("is_member")
	})
})

var AddUserPayload = Type("AddUserPayload", func() {
	Member("name", String, "username", func() {
		Pattern(".+")
		Example("Richard Matthew Stallman")
	})
	Member("password", String, "password", func() {
		Pattern(".+")
		Example("password123")
	})
	Required("name", "password")
})

var ModifyUserPayload = Type("ModifyUserPayload", func() {
	Member("is_register", Boolean, func() {
		Example(true)
	})
	Required("is_register")
})
