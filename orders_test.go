package main

import (
	"testing"

	"time"

	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/app/test"
	"github.com/riku179/regisys-server/models"
)

var orderCtrl = NewOrdersController(service)

func TestAddOrders_NoContent(t *testing.T) {
	r, registerCtx := PrepareUser(Register)
	user, _ := PrepareUser(Normal)
	item := PrepareItems("item1", user.ID)
	defer UserDB.Delete(ctx, user.ID)
	defer UserDB.Delete(ctx, r.ID)
	defer ItemsDB.Delete(ctx, item.ID)

	t.Log("ユーザーの商品を購入する")
	test.AddOrdersNoContent(t, registerCtx, service, orderCtrl, &app.AddOrderPayload{
		ItemID:        item.ID,
		IsMemberPrice: false,
		Quantity:      5,
	})

	var testOrder models.Orders
	err := OrdersDB.Db.Last(&testOrder).Error
	defer OrdersDB.Delete(ctx, testOrder.ID)
	if err == gorm.ErrRecordNotFound {
		t.Fatal("Added Order NOT FOUND in DB")
	}
}

func TestAddOrders_Forbidden(t *testing.T) {
	user, normalCtx := PrepareUser(Normal)
	item := PrepareItems("item1", user.ID)
	defer UserDB.Delete(ctx, user.ID)
	defer ItemsDB.Delete(ctx, item.ID)

	t.Log("Normalユーザーが購入処理をする")
	test.AddOrdersForbidden(t, normalCtx, service, orderCtrl, &app.AddOrderPayload{
		ItemID:        item.ID,
		IsMemberPrice: false,
		Quantity:      5,
	})
}

func TestAddOrders_NotFound(t *testing.T) {
	r, registerCtx := PrepareUser(Register)
	defer UserDB.Delete(ctx, r.ID)

	t.Log("存在しない商品を購入する")
	test.AddOrdersNotFound(t, registerCtx, service, orderCtrl, &app.AddOrderPayload{
		ItemID:        114514,
		IsMemberPrice: false,
		Quantity:      5,
	})
}

func TestDeleteOrders_NoContent(t *testing.T) {
	r, registerCtx := PrepareUser(Register)
	testOrder := PrepareOrders(123)
	defer UserDB.Delete(ctx, r.ID)
	defer OrdersDB.Delete(ctx, testOrder.ID)

	t.Log("オーダーを削除する")
	test.DeleteOrdersNoContent(t, registerCtx, service, orderCtrl, testOrder.ID)

	_, err := OrdersDB.Get(ctx, testOrder.ID)
	if err != gorm.ErrRecordNotFound {
		t.Log("Order has been still remained")
	}
}

func TestDeleteOrders_NotFound(t *testing.T) {
	r, registerCtx := PrepareUser(Register)
	defer UserDB.Delete(ctx, r.ID)

	t.Log("存在しないオーダーを削除する")
	test.DeleteOrdersNotFound(t, registerCtx, service, orderCtrl, 114514)
}

func TestDeleteOrders_Forbidden(t *testing.T) {
	user, normalCtx := PrepareUser(Normal)
	testOrder := PrepareOrders(123)
	defer UserDB.Delete(ctx, user.ID)
	defer UserDB.Delete(ctx, testOrder.ID)

	t.Log("一般ユーザーで削除する")
	test.DeleteOrdersForbidden(t, normalCtx, service, orderCtrl, testOrder.ID)
}

func TestShowOrders_OK(t *testing.T) {
	user, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, user.ID)

	t.Log("全オーダーを取得")
	test.ShowOrdersOK(t, normalCtx, service, orderCtrl, 0, 2147483647, nil)

	t.Log("ある期間のオーダーだけ取得")
	test.ShowOrdersOK(t, normalCtx, service, orderCtrl, 0, int(time.Now().Unix()), nil)

	t.Log("あるユーザーのオーダーだけ取得")
	test.ShowOrdersOK(t, normalCtx, service, orderCtrl, 0, 2147483647, &user.ID)
}

func TestShowOrders_NotFound(t *testing.T) {
	user, normalCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, user.ID)

	t.Log("存在しないユーザーを指定してオーダーを取得")
	id := 114514
	test.ShowOrdersNotFound(t, normalCtx, service, orderCtrl, 0, 2147483647, &id)
}

func PrepareOrders(itemID int) (orders *models.Orders) {
	regiUser := &models.User{Name: Register, Group: Register, IsMember: true}
	UserDB.Add(ctx, regiUser)
	orders = &models.Orders{
		UserID:        regiUser.ID,
		ItemID:        itemID,
		Price:         100,
		Quantity:      5,
		IsMemberPrice: false,
	}
	OrdersDB.Add(ctx, orders)
	return
}
