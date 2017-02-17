package main

import (
	"fmt"
	"os"
	"testing"

	"context"

	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	"github.com/riku179/regisys-server/app"
	"github.com/riku179/regisys-server/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ctx    context.Context = context.Background()
	testPW string          = string(bcrypt.GenerateFromPassword([]byte("foobar"), 10))
)

// ## 目次
// * DBへの接続
// * テスト用テーブルの作成の作成・マイグレーション
// * ミドルウェア・コントローラの登録
// * テスト用データのINSERT
// * +テストの実行+
// * テスト用テーブルのDROP
func TestMain(m *testing.M) {
	// * DBへの接続
	os.Setenv(MYSQL_DATABASE, "testing")
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(db:3306)/%v?parseTime=true&charset=utf8", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))
	if err != nil {
		exitOnFailure(err)
	}
	defer db.Close()
	ItemsDB = models.NewItemsDB(db)
	UserDB = models.NewUserDB(db)
	OrdersDB = models.NewOrdersDB(db)

	// * テスト用テーブルの作成の作成・マイグレーション
	db.CreateTable(&models.Items{}, &models.User{}, &models.Orders{})

	// * ミドルウェア・コントローラの登録
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	jwtMiddleware, err := NewJWTMiddleware()
	exitOnFailure(err)
	app.UseJWTMiddleware(service, jwtMiddleware)
	app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	c := NewItemsController(service)
	app.MountItemsController(service, c)
	c2, err := NewJWTController(service)
	exitOnFailure(err)
	app.MountJWTController(service, c2)
	c3 := NewOrdersController(service)
	app.MountOrdersController(service, c3)
	c4 := NewSwaggerController(service)
	app.MountSwaggerController(service, c4)
	c5 := NewUserController(service)
	app.MountUserController(service, c5)

	// * テスト用データのINSERT
	// 一般権限のMMA部員
	UserDB.Add(ctx, &models.User{
		IsMember: true,
		Name:     "Mr.MMA",
		Password: testPW,
		Group:    Normal,
	})
	//　OB
	UserDB.Add(ctx, &models.User{
		IsMember: false,
		Name:     "Mr.OB",
		Password: testPW,
	})
	// レジ打ち権限のMMA部員
	UserDB.Add(ctx, &models.User{
		IsMember: true,
		Name:     "Mr.Regiser",
		Password: testPW,
		Group:    Register,
	})
	// 管理者権限のMMA部員
	UserDB.Add(ctx, &models.User{
		IsMember: true,
		Name:     "Mr.Admin",
		Password: testPW,
		Group:    Admin,
	})

	// * +テストの実行+
	code := m.Run()

	// * テスト用テーブルのDROP
	db.DropTableIfExists(ItemsDB.TableName(), UserDB.TableName(), OrdersDB.TableName())

	os.Exit(code)
}
