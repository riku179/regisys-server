package main

import "github.com/riku179/regisys-server/models"

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
