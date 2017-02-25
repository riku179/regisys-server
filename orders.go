package main

import (
	"context"

	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/models"
	"github.com/riku179/regisys-server/user"
)

// OrdersController implements the orders resource.
type OrdersController struct {
	*goa.Controller
	DB *models.OrdersDB
}

// NewOrdersController creates a orders controller.
func NewOrdersController(service *goa.Service) *OrdersController {
	return &OrdersController{
		Controller: service.NewController("OrdersController"),
	}
}

// Add runs the add action.
func (c *OrdersController) Add(ctx *app.AddOrdersContext) error {
	reqUser := user.FromContext(ctx)
	// Register以上の権限が無いとダメ
	if reqUser.Group != Register && reqUser.Group != Admin {
		return ctx.Forbidden()
	}

	item, err := ItemsDB.Get(ctx, ctx.Payload.ItemID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}

	order := models.OrdersFromAddOrderPayload(ctx.Payload)
	order.UserID = reqUser.ID // レジ担当者ID

	// メンバー価格の適用
	if ctx.Payload.IsMemberPrice {
		order.Price = item.MemberPrice
	} else {
		order.Price = item.Price
	}

	OrdersDB.Add(ctx, order)

	return ctx.NoContent()
}

// Delete runs the delete action.
func (c *OrdersController) Delete(ctx *app.DeleteOrdersContext) error {
	reqUser := user.FromContext(ctx)
	if reqUser.Group != Register && reqUser.Group != Admin {
		return ctx.Forbidden()
	}

	_, err := ItemsDB.Get(ctx, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		// 指定された商品が存在しない
		return ctx.NotFound()
	}

	OrdersDB.Delete(ctx, ctx.ID)

	return ctx.NoContent()
}

// Show runs the show action.
func (c *OrdersController) Show(ctx *app.ShowOrdersContext) error {
	timeStart := time.Unix(int64(ctx.TimeStart), 0)
	timeEnd := time.Unix(int64(ctx.TimeEnd), 0)

	if ctx.User == nil {
		// 全ユーザー
		orders := []*models.Orders{}
		OrdersDB.Db.Where(
			"created_at >= ? AND created_at <= ?", timeStart, timeEnd,
		).Find(&orders)
		res := OrdersToRegisysOrders(ctx, orders)

		return ctx.OK(res)

	} else {
		// 指定ユーザーの取引のみ
		orders := []*models.Orders{}

		OrdersDB.Db.Where(
			"user_id = ?", ctx.User,
		).Where(
			"created_at >= ? AND created_at <= ?", timeStart, timeEnd,
		).Find(&orders)

		res := OrdersToRegisysOrders(ctx, orders)

		return ctx.OK(res)
	}

	//return ctx.OK(res)
	return ctx.NotFound()
}

// []*models.Orders -> []*app.RegisysOrders
func OrdersToRegisysOrders(ctx context.Context, orders []*models.Orders) app.RegisysOrdersCollection {
	res := app.RegisysOrdersCollection{}
	for _, v := range orders {
		item, _ := ItemsDB.Get(ctx, v.ItemID)
		res = append(res, &app.RegisysOrders{
			ID:       v.ID,
			ItemID:   v.ItemID,
			Price:    v.Price,
			Quantity: v.Quantity,
			UserID:   v.UserID,
			Datetime: int(v.CreatedAt.Unix()),
			ItemName: item.ItemName,
		})
	}
	return res
}
