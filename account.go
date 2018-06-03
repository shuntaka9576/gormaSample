package main

import (
	"github.com/goadesign/goa"
	"github.com/shuntaka9576/gormaSample/app"
	"github.com/shuntaka9576/gormaSample/models"
	"github.com/jinzhu/gorm"
)

// AccountController implements the account resource.
type AccountsController struct {
	*goa.Controller
	db *gorm.DB
}

// NewAccountsController creates a accounts controller.
func NewAccountsController(service *goa.Service, db *gorm.DB) *AccountsController {
	return &AccountsController{
		Controller: service.NewController("AccountsController"),
		db:         db,
	}
}

// Show runs the show action.
func (c *AccountsController) Show(ctx *app.ShowAccountContext) error {
	// AccountsController_Show: start_implement

	// Put your logic here
	adb := models.NewAccountDB(c.db)
	a, err := adb.OneAccount(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	// AccountsController_Show: end_implement
	res := &app.Account{}
	res = a
	return ctx.OK(res)
}

