package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("orders", func() {
	Description("Provide orders")
	BasePath("/orders")

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("show", func() {
		Description("Get orders")
		Routing(GET(""))
		Params(func() {
			Param("user", String, "Unique user ID")
			Param("date", DateTime, "Order date")
		})

		Response(OK, func() {
			Media(CollectionOf(OrderMedia))
		})
		Response(NotFound)
	})

	Action("add", func() {
		Description("Add order")
		Routing(POST(""))
		Payload(AddOrderPayload)

		Response(NoContent)
		Response(Forbidden)
	})

	Action("delete", func() {
		Description("Disable order")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "Order ID")
		})

		Response(NoContent)
		Response(NotFound)
		Response(Forbidden)
	})
})

var OrderMedia = MediaType("application/vnd.goa.example.orders+json", func() {
	Description("An order")
	Attributes(func() {
		Attribute("id", Integer, "Unique order ID")
		Attribute("item_id", Integer, "Unique item ID")
		Attribute("item_name", String, "item name")
		Attribute("quantity", Integer, "item quantity")
		Attribute("price", Integer, "item price")
		Attribute("user_id", Integer, "Register's user ID")
		Attribute("date", DateTime, "Order date")
		Required("id", "item_id", "item_name", "quantity", "price", "user_id", "date")
	})
	View("default", func() {
		Attribute("id")
		Attribute("item_id")
		Attribute("item_name")
		Attribute("quantity")
		Attribute("price")
		Attribute("user_id")
		Attribute("date")
	})
})

var AddOrderPayload = Type("AddOrderPayload", func() {
	Member("item_id", Integer, "Unique item ID")
	Member("quantity", Integer, "item quantity")
	Member("price", Integer, "item price")
	Member("user_id", Integer, "Register's user ID")
	Required("item_id", "quantity", "price", "user_id")
})
