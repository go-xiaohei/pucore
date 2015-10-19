package user

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
	"github.com/go-xiaohei/pucore/utils"
)

type Module struct{}

func (um *Module) Id() string {
	return "USER"
}

func (um *Module) Prepare(ctx *pucore.ModuleContext) error {
	return nil
}

func (um *Module) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (um *Module) Disable(ctx *pucore.ModuleContext) error {
	return pucore.ErrModuleDisableIgnore
}

func (um *Module) Install(ctx *pucore.ModuleContext) error {
	database := new(app.Db)
	ctx.Injector.Get(database)
	var err error
	if err = database.Sync2(new(User)); err != nil {
		return err
	}

	// insert default user
	user := &User{
		Id:        9,
		Name:      "admin",
		Email:     "admin@example.com",
		Nick:      "admin",
		Profile:   "this is administrator",
		Role:      USER_ROLE_ADMIN,
		Status:    USER_STATUS_ACTIVE,
		AvatarUrl: utils.Gravatar("admin@admin.com"),
	}
	user.SetPassword("123456789")
	if _, err := database.Insert(user); err != nil {
		return err
	}
	return nil
}
