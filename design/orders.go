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
		Attribute("item", String, "Item name")
		Attribute("quantity", Integer, "Item quantity")
		Attribute("price", Integer, "Item price")
		Attribute("register", String, "Register's name")
		Attribute("date", DateTime, "Order date")
		Attribute("available", Boolean, "Is available order")
		Required("id", "item_id", "item", "quantity", "price", "register", "date", "available")
	})
	View("default", func() {
		Attribute("id")
		Attribute("item_id")
		Attribute("item")
		Attribute("quantity")
		Attribute("price")
		Attribute("register")
		Attribute("date")
		Attribute("available")
	})
})

var AddOrderPayload = Type("AddOrderPayload", func() {
	Member("item_id", Integer, "Unique item ID")
	Member("quantity", Integer, "Item quantity")
	Member("price", Integer, "Item price")
	Member("register", String, "Register's name")
	Member("date", DateTime, "Order date")
	Required("item_id", "quantity", "price", "register", "date")
})
