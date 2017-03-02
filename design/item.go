package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("items", func() {
	Description("Provide items")
	BasePath("/item")

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("show", func() {
		Description("Get items")
		Routing(GET(""))
		Params(func() {
			Param("user", Integer, "User ID", func() {
				Example(1)
			})
		})

		Response(OK, func() {
			Media(CollectionOf(ItemMedia))
		})
		Response(NotFound)
	})

	Action("add", func() {
		Description("Add item")
		Routing(POST(""))

		Payload(AddItemPayload)

		Response(NoContent)
	})

	Action("modify", func() {
		Description("Modify item")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id", Integer, "Unique item ID")
		})
		Payload(ModifyItemPayload)

		Response(NoContent)
		Response(NotFound)
		Response(Forbidden)
	})

	Action("delete", func() {
		Description("Delete item")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "Unique item ID")
		})

		Response(NoContent)
		Response(NotFound)
		Response(Forbidden)
	})
})

var ItemMedia = MediaType("application/vnd.regisys.items+json", func() {
	Description("An item")

	Attributes(func() {
		Attribute("id", Integer, "Unique item ID", func() {
			Example(1)
		})
		Attribute("item_name", String, "item name", func() {
			Example("Thinkpad X1 Carbon")
		})
		Attribute("price", Integer, "item price", func() {
			Example(120000)
		})
		Attribute("member_price", Integer, "Membership discount", func() {
			Example(100000)
		})
		Attribute("user_id", Integer, "Unique User ID", func() {
			Example(1001)
		})
		Attribute("quantity", Integer, "item quantity", func() {
			Example(1)
		})
		Required("id", "item_name", "price", "member_price", "user_id", "quantity")
	})

	View("default", func() {
		Attribute("id")
		Attribute("item_name")
		Attribute("price")
		Attribute("member_price")
		Attribute("user_id")
		Attribute("quantity")
	})
})

var AddItemPayload = Type("AddItemPayload", func() {
	Member("item_name", String, "item name", func() {
		Pattern(".+")
		Example("DDR3 RAM 2G")
	})
	Member("price", Integer, "item price", func() {
		Minimum(0)
		Example(1000)
	})
	Member("member_price", Integer, "Membership discount", func() {
		Minimum(0)
		Example(800)
	})
	Member("quantity", Integer, "item quantity", func() {
		Minimum(1)
		Example(4)
	})
	Required("item_name", "price", "member_price", "quantity")
})

var ModifyItemPayload = Type("ModifyItemPayload", func() {
	Member("item_name", String, "item name", func() {
		Pattern(".+")
		Example("Mac Book Air")
	})
	Member("price", Integer, "item price", func() {
		Minimum(0)
		Example(500)
	})
	Member("member_price", Integer, "Membership discount", func() {
		Minimum(0)
		Example(400)
	})
	Member("quantity", Integer, "item quantity", func() {
		Minimum(1)
		Example(6)
	})
})
