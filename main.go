//go:generate goagen bootstrap -d github.com/shuntaka9576/gormaSample/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/shuntaka9576/gormaSample/app"
	"flag"
	"log"
	"github.com/shuntaka9576/gormaSample/database"
	"github.com/shuntaka9576/gormaSample/controller"
)

func main() {
	// Create service
	service := goa.New("cellar")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// *****************
	var (
		env   = flag.String("env", "development", "application envirionment (production, development etc.)")
		dbrun = flag.Bool("dbrun", false, "database run mode")
	)
	flag.Parse()

	if *dbrun {
		log.Print("DB読み込み処理テスト")
		cs, err := database.NewConfigsFromFile("dbconfig.yml")
		if err != nil {
			log.Fatalf("cannot open database configuration. exit. %s", err)
		}
		dbcon, err := cs.Open(*env)
		if err != nil {
			log.Fatalf("database initialization failed: %s", err)
		}
		// Mount "accounts" controller
		a := controller.NewAccountsController(service, dbcon)
		app.MountAccountController(service, a)
	}
	// *****************

	// Mount "account" controller
	//c := NewAccountController(service)
	//app.MountAccountController(service, c)
	//// Mount "bottle" controller
	//c2 := NewBottleController(service)
	//app.MountBottleController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
