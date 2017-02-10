package main

import (
	"fmt"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys/app"
	"github.com/riku179/regisys/models"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv(MYSQL_DATABASE, "testing")
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(db:3306)/%v?parseTime=true&charset=utf8", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))
	if err != nil {
		exitOnFailure(err)
	}
	defer db.Close()
	ItemsDB = models.NewItemsDB(db)
	UserDB = models.NewUserDB(db)
	OrdersDB = models.NewOrdersDB(db)

	//// ################# DROP TABLE when debug mode ###################
	if DEBUG == "TRUE" {
		db.DropTableIfExists(ItemsDB.TableName(), UserDB.TableName(), OrdersDB.TableName())
	}
	//// ################################################################

	//db.AutoMigrate(&models.Goods{}, &models.User{}, &models.Orders{})
	db.CreateTable(&models.Items{}, &models.User{}, &models.Orders{})

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middlewares
	jwtMiddleware, err := NewJWTMiddleware()
	exitOnFailure(err)
	app.UseJWTMiddleware(service, jwtMiddleware)
	// Security middleware used to secure the creation of JWT tokens.
	app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	// Mount "goods" controller
	c := NewItemsController(service)
	app.MountItemsController(service, c)
	// Mount "jwt" controller
	c2, err := NewJWTController(service)
	exitOnFailure(err)
	app.MountJWTController(service, c2)
	// Mount "orders" controller
	c3 := NewOrdersController(service)
	app.MountOrdersController(service, c3)
	// Mount "swagger" controller
	c4 := NewSwaggerController(service)
	app.MountSwaggerController(service, c4)
	// Mount "user" controller
	c5 := NewUserController(service)
	app.MountUserController(service, c5)

	code := m.Run()
	os.Exit(code)
}
