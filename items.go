package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/models"
	"github.com/riku179/regisys-server/user"
)

// ItemsController implements the items resource.
type ItemsController struct {
	*goa.Controller
	DB *models.ItemsDB
}

// NewItemsController creates a items controller.
func NewItemsController(service *goa.Service) *ItemsController {
	return &ItemsController{
		Controller: service.NewController("ItemsController"),
	}
}

// Add runs the add action.
func (c *ItemsController) Add(ctx *app.AddItemsContext) error {
	reqUser := user.FromContext(ctx)

	items := models.ItemsFromAddItemPayload(ctx.Payload)
	items.UserID = reqUser.ID // 商品の所有ユーザのID

	ItemsDB.Add(ctx, items)

	return ctx.NoContent()
}

// Delete runs the delete action.
func (c *ItemsController) Delete(ctx *app.DeleteItemsContext) error {
	reqUser := user.FromContext(ctx)

	reqItem, err := ItemsDB.Get(ctx, ctx.ID)

	if err == gorm.ErrRecordNotFound {
		// 指定された商品が無い
		return ctx.NotFound()
	}
	if reqItem.UserID != reqUser.ID {
		// 指定された商品の所有者IDがリクエストしたユーザーののIDと一致しない
		return ctx.Forbidden()
	}

	var relatedOrder []*models.Orders
	OrdersDB.Db.Where("item_id = ?", ctx.ID).First(&relatedOrder)
	if len(relatedOrder) != 0 {
		// 既に購入処理がされた商品は削除できない
		return ctx.Forbidden()
	} else {
		// OK
		ItemsDB.Delete(ctx, ctx.ID)
	}
	return ctx.NoContent()
}

// Modify runs the modify action.
func (c *ItemsController) Modify(ctx *app.ModifyItemsContext) error {
	reqUser := user.FromContext(ctx)

	reqItem, err := ItemsDB.Get(ctx, ctx.ID)

	if err == gorm.ErrRecordNotFound {
		// 指定された商品が無い
		return ctx.NotFound()
	}
	if reqItem.UserID != reqUser.ID {
		// 指定された商品の所有者IDがリクエストしたユーザーののIDと一致しない
		return ctx.Forbidden()
	}

	var relatedOrder []*models.Orders
	OrdersDB.Db.Where(models.Orders{ItemID: ctx.ID}).First(&relatedOrder)

	if len(relatedOrder) != 0 {
		// 既に購入処理がされた商品は削除できない
		if ctx.Payload.ItemName == nil && ctx.Payload.Quantity == nil {
			// が、価格の変更だけならOK
			ItemsDB.UpdateFromModifyItemPayload(ctx, ctx.Payload, ctx.ID)
			return ctx.NoContent()
		}
		return ctx.Forbidden()
	} else {
		// OK
		ItemsDB.UpdateFromModifyItemPayload(ctx, ctx.Payload, ctx.ID)
	}
	return ctx.NoContent()
}

// Show runs the show action.
func (c *ItemsController) Show(ctx *app.ShowItemsContext) error {
	var res app.RegisysItemsCollection

	// ユーザーが指定されていない場合は全てのItemを返す
	if ctx.User == nil {
		return ctx.OK(ItemsDB.ListRegisysItems(ctx))
	}

	_, err := UserDB.Get(ctx, *ctx.User)
	if err == gorm.ErrRecordNotFound {
		// ユーザーが存在しない
		return ctx.NotFound()
	} else {
		// ユーザーが指定されている場合はそのユーザーのItemのみ返す
		var items []*models.Items
		ItemsDB.Db.Where("user_id = ?", ctx.User).Find(&items)

		for _, t := range items {
			res = append(res, t.ItemsToRegisysItems())
		}
	}

	return ctx.OK(res)
}
