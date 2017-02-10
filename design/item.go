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
			Param("user", Integer, "User ID")
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
		Attribute("id", Integer, "Unique item ID")
		Attribute("item_name", String, "item name")
		Attribute("price", Integer, "item price")
		Attribute("member_price", Integer, "Membership discount")
		Attribute("user_id", Integer, "Unique User ID")
		Attribute("user_name", String, "Username")
		Attribute("quantity", Integer, "item quantity")
		Required("id", "item_name", "price", "member_price", "user_id", "user_name", "quantity")
	})

	View("default", func() {
		Attribute("id")
		Attribute("item_name")
		Attribute("price")
		Attribute("member_price")
		Attribute("user_id")
		Attribute("user_name")
		Attribute("quantity")
	})
})

var AddItemPayload = Type("AddGoodsPayload", func() {
	Member("item_name", String, "item name")
	Member("price", Integer, "item price")
	Member("member_price", Integer, "Membership discount")
	Member("quantity", Integer, "item quantity")
	Required("item_name", "price", "member_price", "quantity")
})

var ModifyItemPayload = Type("ModifyGoodsPayload", func() {
	Member("id", Integer, "Unique item ID")
	Member("item_name", String, "item name")
	Member("price", Integer, "item price")
	Member("member_price", Integer, "Membership discount")
	Member("quantity", Integer, "item quantity")
	Required("id")
})
