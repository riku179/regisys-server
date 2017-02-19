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
)

var (
	ctx    context.Context = context.Background()
	testPW                 = "foobar" // password for all testUser
	//testUser                 = map[string]string{
	//	"MMA":  "Mr.MMA",       // user of MMA member in 'Normal' group
	//	"MMA2": "Mr.MMA2",      // second person in same group
	//	"OB":   "Mr.OB",        // OB user(not MMA member) in 'Normal' group
	//	"OB2":  "Mr.OB2",       // second person in same group
	//	"Reg":  "Mr.Register",  // user of 'Register' group
	//	"Reg2": "Mr.Register2", // second person in same group
	//	"Adm":  "Mr.Admin",     // user of 'Admin' group
	//	"Adm2": "Mr.Admin2",    // second person in same group
	//}
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
	insertTestUser()

	// * +テストの実行+
	code := m.Run()

	// * テスト用テーブルのDROP
	db.DropTableIfExists(ItemsDB.TableName(), UserDB.TableName(), OrdersDB.TableName())

	os.Exit(code)
}

func insertTestUser() {
	//bPW, _ := bcrypt.GenerateFromPassword([]byte(testPW), 10)
	//pw := string(bPW)

	////一般権限のMMA部員
	//UserDB.Add(ctx, &models.User{
	//	IsMember: true,
	//	Name:     testUser["MMA"],
	//	Password: pw,
	//	Group:    Normal,
	//})
	//UserDB.Add(ctx, &models.User{
	//	IsMember: true,
	//	Name:     testUser["MMA2"],
	//	Password: pw,
	//	Group:    Normal,
	//})
	////　OB
	//UserDB.Add(ctx, &models.User{
	//	IsMember: false,
	//	Name:     testUser["OB"],
	//	Password: pw,
	//})
	//UserDB.Add(ctx, &models.User{
	//	IsMember: false,
	//	Name:     testUser["OB2"],
	//	Password: pw,
	//})
	//// レジ打ち権限のMMA部員
	//UserDB.Add(ctx, &models.User{
	//	IsMember: true,
	//	Name:     testUser["Reg"],
	//	Password: pw,
	//	Group:    Register,
	//})
	//UserDB.Add(ctx, &models.User{
	//	IsMember: true,
	//	Name:     testUser["Reg2"],
	//	Password: pw,
	//	Group:    Register,
	//})
	//// 管理者権限のMMA部員
	//UserDB.Add(ctx, &models.User{
	//	IsMember: true,
	//	Name:     testUser["Adm"],
	//	Password: pw,
	//	Group:    Admin,
	//})
	//UserDB.Add(ctx, &models.User{
	//	IsMember: true,
	//	Name:     testUser["Adm2"],
	//	Password: pw,
	//	Group:    Admin,
	//})
}
