package auth

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
	"github.com/go-xiaohei/pucore/module/user"
	"sync"
)

type Module struct {
	isEnable bool
	mutex    sync.Mutex
}

func (am *Module) Id() string {
	return "AUTHORIZE"
}

func (am *Module) Prepare(ctx *pucore.ModuleContext) error {
	if !ctx.Modular.Has(new(user.Module)) {
		return pucore.ErrModuleDependsOn(am, new(user.Module))
	}
	return nil
}

func (am *Module) Enable(ctx *pucore.ModuleContext) error {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	am.isEnable = true
	return nil
}

func (am *Module) Disable(ctx *pucore.ModuleContext) error {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	am.isEnable = false
	return nil
}

func (am *Module) Install(ctx *pucore.ModuleContext) error {
	database := new(app.Db)
	ctx.Injector.Get(database)
	var err error
	if err = database.Sync2(new(Token)); err != nil {
		return err
	}
	return nil
}
