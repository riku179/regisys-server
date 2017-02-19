//go:generate goagen bootstrap -d regisys/design

package main

import (
	"fmt"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/models"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	errValidationFailed = goa.NewErrorClass("validation_failed", 401)
	ItemsDB             *models.ItemsDB
	UserDB              *models.UserDB
	OrdersDB            *models.OrdersDB
	MYSQL_USER          = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD      = os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE      = os.Getenv("MYSQL_DATABASE")
	DEBUG               = os.Getenv("DEBUG")
	service             = goa.New("regisys")
)

const (
	Admin    = "admin"
	Register = "register"
	Normal   = "normal"
)

func main() {
	// Create service

	// Connect DB
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
