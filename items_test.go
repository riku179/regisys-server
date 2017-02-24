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

	test.DeleteItemsNoContent(t, userCtx, service, itemCtrl, testItem.ID)

	if testItem.DeletedAt != nil {
		t.Fatal("Speficied item didn't deleted!")
	}
}

func TestDeleteItem_NotFound(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)
	// ID: 114514 doesn't exist
	test.DeleteItemsNotFound(t, userCtx, service, itemCtrl, 114514)
}

func TestDeleteItem_Forbidden(t *testing.T) {
	normalUser, userCtx := PrepareUser(Normal)
	adminUser, _ := PrepareUser(Admin)
	defer UserDB.Delete(ctx, normalUser.ID)
	defer UserDB.Delete(ctx, adminUser.ID)

	adminItem := PrepareItems("foo", adminUser.ID)
	defer ItemsDB.Delete(ctx, adminItem.ID)

	// usr cannot delete other's items
	test.DeleteItemsForbidden(t, userCtx, service, itemCtrl, adminItem.ID)
	//TODO 既に購入処理がされていた場合、商品の編集が拒否されるテスト
}

func TestModifyItem_NoContent(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)

	testItem := PrepareItems("foo", reqUser.ID)
	defer ItemsDB.Delete(ctx, testItem.ID)

	newName := "bar"
	test.ModifyItemsNoContent(
		t, userCtx, service, itemCtrl, testItem.ID, &app.ModifyItemPayload{
			ItemName: &newName,
		})
	//TODO 既に会計情報がついていても値段の変更のみするテスト

	testItem, _ = ItemsDB.Get(ctx, testItem.ID)
	if testItem.ItemName != "bar" {
		t.Fatalf("Expected ItemName: bar, but receive: %+v", testItem.ItemName)
	}

}

func TestModifyItem_NotFound(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)
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

	newName := "bar"
	test.ModifyItemsForbidden(t, userCtx, service, itemCtrl, adminItem.ID, &app.ModifyItemPayload{
		ItemName: &newName,
	})
	testOrder := &models.Orders{
		ItemID:        adminItem.ID, //10
		Quantity:      1,
		IsMemberPrice: false,
		Price:         100,
		UserID:        adminUser.ID,
	}
	OrdersDB.Add(ctx, testOrder)
	defer OrdersDB.Delete(ctx, testOrder.ID)

	test.ModifyItemsForbidden(t, adminCtx, service, itemCtrl, adminItem.ID, &app.ModifyItemPayload{
		ItemName: &newName,
	})
}

func TestShowItem_OK(t *testing.T) {
	reqUser, userCtx := PrepareUser(Normal)
	defer UserDB.Delete(ctx, reqUser.ID)

	testItem := PrepareItems("foo", reqUser.ID)
	defer ItemsDB.Delete(ctx, testItem.ID)

	test.ShowItemsOK(t, userCtx, service, itemCtrl, &reqUser.ID)
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
