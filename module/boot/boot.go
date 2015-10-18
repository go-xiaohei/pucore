package boot

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
)

type Boot struct {
}

func (b *Boot) Id() string {
	return "BOOT"
}

func (b *Boot) Prepare(ctx *pucore.ModuleContext) error {
	var (
		config   = new(app.Config)
		database = new(app.Db)
	)

	// read config file or create file, try to
	ctx.Injector.Get(config)
	if err := config.Sync(); err != nil {
		return err
	}
	ctx.Injector.Set(config)
	if config.IsNew {
		flag := InstallFlag(true)
		ctx.Injector.Set(&flag)
	}

	// connect to database
	ctx.Injector.Get(database)
	database.Driver = config.Db.Driver
	database.DSN = config.Db.DSN
	if err := database.Connect(); err != nil {
		return err
	}
	ctx.Injector.Set(database)

	return nil

}

func (b *Boot) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (b *Boot) Disable(ctx *pucore.ModuleContext) error {
	return nil
}
