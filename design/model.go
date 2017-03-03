package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("RegisysStorageGroup", func() {
	Description("Global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("MySQL relational store")
		Model("Items", func() {
			BuildsFrom(func() {
				Payload("items", "add")
				Payload("items", "modify")
			})
			RendersTo(ItemMedia)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("Unique item ID")
			})
			Field("item_name", gorma.String, func() {
				SQLTag("not null;default:false")
			})
			Field("price", gorma.Integer, func() {
				SQLTag("not null;default:false")
			})
			Field("member_price", gorma.Integer, func() {
				SQLTag("not null;default:false")
			})
			Field("quantity", gorma.Integer, func() {
				SQLTag("not null;default:false")
			})

			Field("CreatedAt", gorma.Timestamp)
			Field("UpdatedAt", gorma.Timestamp)
			Field("DeletedAt", gorma.Timestamp)

			HasMany("Orders", "Orders")
		})

		Model("User", func() {
			RendersTo(ShowUserMedia)
			BuildsFrom(func() {
				Payload("user", "add")
				Payload("user", "modify")
			})
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("Unique user ID")
			})
			Field("name", gorma.String, func() {
				SQLTag("not null;default:false")
			})
			Field("is_register", gorma.Boolean, func() {
				SQLTag("not null;default:false")
			})
			Field("is_member", gorma.Boolean, func() {
				SQLTag("not null;default:false")
			})

			Field("CreatedAt", gorma.Timestamp)
			Field("UpdatedAt", gorma.Timestamp)
			Field("DeletedAt", gorma.Timestamp)

			HasMany("Items", "Items")
			HasMany("RegisteredOrders", "Orders")
		})

		Model("Orders", func() {
			BuildsFrom(func() {
				Payload("orders", "add")
			})
			RendersTo(OrderMedia)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("Unique order ID")
			})
			Field("item_id", gorma.Integer, func() {
				SQLTag("not null;default:false")
			})
			Field("price", gorma.Integer, func() {
				SQLTag("not null;default:false")
			})
			Field("quantity", gorma.Integer, func() {
				SQLTag("not null;default:false")
			})

			Field("CreatedAt", gorma.Timestamp)
			Field("DeletedAt", gorma.Timestamp)
		})
	})
})
