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
			Param("user", Integer, "Unique user ID", func() {
				Minimum(0)
			})
			Param("time_start", Integer, "Start Order date(UnixTime)", func() {
				Minimum(0)
				Default(0)
			})
			Param("time_end", Integer, "End Order date(UnixTime)", func() {
				Minimum(0)
				Maximum(2147483647)
				Default(2147483647)
			})
			Required("time_start", "time_end")
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
		Response(NotFound)
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

var OrderMedia = MediaType("application/vnd.regisys.orders+json", func() {
	Description("An order")
	Attributes(func() {
		Attribute("id", Integer, "Unique order ID")
		Attribute("item_id", Integer, "Unique item ID")
		Attribute("item_name", String, "item name")
		Attribute("quantity", Integer, "item quantity")
		Attribute("price", Integer, "item price")
		Attribute("user_id", Integer, "Register's user ID")
		Attribute("datetime", Integer, "Order datetime")
		Required("id", "item_id", "item_name", "quantity", "price", "user_id", "datetime")
	})
	View("default", func() {
		Attribute("id")
		Attribute("item_id")
		Attribute("item_name")
		Attribute("quantity")
		Attribute("price")
		Attribute("user_id")
		Attribute("datetime")
	})
})

var AddOrderPayload = Type("AddOrderPayload", func() {
	Member("item_id", Integer, "Unique item ID")
	Member("quantity", Integer, "item quantity", func() {
		Minimum(1)
	})
	Member("is_member_price", Boolean, "Is it bought for member's price")
	Required("item_id", "quantity", "is_member_price")
})
