package main

import (
	"testing"

	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/app/test"
	"github.com/riku179/regisys-server/models"
)

var itemCtrl = NewItemsController(service)

func TestAddItem_NoContent(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)

	t.Log("商品を追加")
	test.AddItemsNoContent(t, userCtx, service, itemCtrl, &app.AddItemPayload{
		ItemName:    "foo",
		Quantity:    3,
		Price:       1000,
		MemberPrice: 800,
	})

	var testItem models.Items
	ItemsDB.Db.Last(&testItem)

	defer ItemsDB.Delete(ctx, testItem.ID)
	if testItem.ItemName != "foo" {
		t.Fatalf("Expected ItemName: hoge, but received: %+v", testItem.ItemName)
	} else if reqUser.ID != testItem.UserID {
		t.Fatalf("Expected UserID: %+v, but received: %+v", reqUser.ID, testItem.UserID)
	}
}

func TestDeleteItem_NoContent(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)

	testItem := PrepareItems("foo", reqUser.ID)
	defer ItemsDB.Delete(ctx, testItem.ID)

	t.Log("まだ購入処理のされていない自分の商品を削除")
	test.DeleteItemsNoContent(t, userCtx, service, itemCtrl, testItem.ID)

	if testItem.DeletedAt != nil {
		t.Fatal("Speficied item didn't deleted!")
	}
}

func TestDeleteItem_NotFound(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)

	t.Log("存在しない商品を削除")
	test.DeleteItemsNotFound(t, userCtx, service, itemCtrl, 114514)
}

func TestDeleteItem_Forbidden(t *testing.T) {
	normalUser, userCtx := PrepareUser(Normal)
	adminUser, _ := PrepareUser(Admin)
	defer UserDB.Delete(ctx, normalUser.ID)
	defer UserDB.Delete(ctx, adminUser.ID)

	t.Log("他人の商品を削除")
	adminItem := PrepareItems("foo", adminUser.ID)
	defer ItemsDB.Delete(ctx, adminItem.ID)

	test.DeleteItemsForbidden(t, userCtx, service, itemCtrl, adminItem.ID)

	t.Log("既に購入処理がされている自分の商品を削除")
	testItem := PrepareItems("hoge", normalUser.ID)
	testOrder := PrepareOrders(testItem.ID)
	defer ItemsDB.Delete(ctx, testItem.ID)
	defer OrdersDB.Delete(ctx, testOrder.ID)

	test.DeleteItemsForbidden(
		t, userCtx, service, itemCtrl, testItem.ID)
}

func TestModifyItem_NoContent(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)
	testItem := PrepareItems("foo", reqUser.ID)
	defer ItemsDB.Delete(ctx, testItem.ID)

	t.Log("まだ購入処理のされていない自分の商品の名前を更新")
	newName := "bar"
	test.ModifyItemsNoContent(
		t, userCtx, service, itemCtrl, testItem.ID, &app.ModifyItemPayload{
			ItemName: &newName,
		})

	testItem, _ = ItemsDB.Get(ctx, testItem.ID)
	if testItem.ItemName != "bar" {
		t.Fatalf("Expected ItemName: bar, but receive: %+v", testItem.ItemName)
	}

	t.Log("既に購入情報がついている自分の商品の値段のみ更新")
	testItem = PrepareItems("hoge", reqUser.ID)
	testOrder := PrepareOrders(testItem.ID)
	defer ItemsDB.Delete(ctx, testItem.ID)
	defer OrdersDB.Delete(ctx, testOrder.ID)

	newPrice := 114514
	test.ModifyItemsNoContent(
		t, userCtx, service, itemCtrl, testItem.ID, &app.ModifyItemPayload{
			Price: &newPrice,
		})

	testItem, _ = ItemsDB.Get(ctx, testItem.ID)
	if testItem.Price != newPrice {
		t.Fatalf("Expected Price: 114514, but receive: %+v", testItem.Price)
	}
}

func TestModifyItem_NotFound(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)

	t.Log("存在しない自分の商品を更新")
	newName := "bar"
	test.ModifyItemsNotFound(
		t, userCtx, service, itemCtrl, 114514, &app.ModifyItemPayload{
			ItemName: &newName,
		})
}

func TestModifyItem_Forbidden(t *testing.T) {
	normalUser, userCtx := PrepareUser(Normal)
	adminUser, adminCtx := PrepareUser(Admin)
	defer UserDB.Delete(ctx, normalUser.ID)
	defer UserDB.Delete(ctx, adminUser.ID)

	adminItem := PrepareItems("foo", adminUser.ID)
	defer ItemsDB.Delete(ctx, adminItem.ID)

	t.Log("他人の商品を更新")
	newName := "bar"
	test.ModifyItemsForbidden(t, userCtx, service, itemCtrl, adminItem.ID, &app.ModifyItemPayload{
		ItemName: &newName,
	})

	testOrder := PrepareOrders(adminItem.ID)
	defer OrdersDB.Delete(ctx, testOrder.ID)

	t.Log("既に購入処理が存在する自分の商品の価格以外を更新")
	test.ModifyItemsForbidden(t, adminCtx, service, itemCtrl, adminItem.ID, &app.ModifyItemPayload{
		ItemName: &newName,
	})
	newQuantitiy := 114514
	test.ModifyItemsForbidden(t, adminCtx, service, itemCtrl, adminItem.ID, &app.ModifyItemPayload{
		Quantity: &newQuantitiy,
	})

}

func TestShowItem_OK(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)

	testItem := PrepareItems("foo", reqUser.ID)
	defer ItemsDB.Delete(ctx, testItem.ID)

	t.Log("自分の商品を表示")
	test.ShowItemsOK(t, userCtx, service, itemCtrl, &reqUser.ID)

	t.Log("全ての商品を表示")
	// UserIDが指定されない場合は全てのItemを返す
	test.ShowItemsOK(t, userCtx, service, itemCtrl, nil)
}

func PrepareItems(name string, userID int) (item *models.Items) {
	item = &models.Items{
		ItemName:    name,
		Quantity:    3,
		Price:       100,
		MemberPrice: 90,
		UserID:      userID,
	}
	ItemsDB.Add(ctx, item)
	return
}
