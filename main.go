//go:generate goagen bootstrap -d regisys/design

package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/riku179/regisys/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"github.com/riku179/regisys/models"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
)

func main() {
	// Create service
	service := goa.New("regisys")

	// Connect DB
	db, err := gorm.Open("mysql", "admin:foobar@tcp(db:3306)/regisys?charset=utf8")
	if err != nil {
		exitOnFailure(err)
	}
	defer db.Close()
	ItemDB := models.NewItemsDB(db)
	UserDB := models.NewUserDB(db)
	OrderDB := models.NewOrdersDB(db)

	//// ################ Only for Develop Environment ##################
	db.DropTableIfExists(ItemDB.TableName(), UserDB.TableName(), OrderDB.TableName())
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

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

// exitOnFailure prints a fatal error message and exits the process with status 1.
func exitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
	os.Exit(1)
}
