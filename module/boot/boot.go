package boot

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
	"gopkg.in/inconshreveable/log15.v2"
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
		server   = new(app.Web)
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
	database.ShowSQL = true
	ctx.Injector.Set(database)
	log15.Debug("Boot.Database.Open." + database.Driver)

	// start web server
	ctx.Injector.Get(server)
	server.Host = config.Http.Host
	server.Port = config.Http.Port
	server.Protocol = config.Http.Protocol
	if err := server.Listen(); err != nil {
		return err
	}
	ctx.Injector.Set(server)
	log15.Debug("Boot.Server.Open." + server.Host + ":" + server.Port)

	// add close hook to modular
	ctx.Modular.Hook(pucore.MODULAR_HOOK_ALL_AFTER, func(_ pucore.IModule) error {
		/*var (
			database = new(app.Db)
			server   = new(app.Web)
		)*/
		ctx.Injector.Get(database, server) // we need read it again
		if err := server.Close(); err != nil {
			log15.Warn("Boot.Server.Close", "error", err)
		}
		log15.Debug("Boot.Server.Close")
		if err := database.Close(); err != nil {
			log15.Warn("Boot.Database.Close", "error", err)
		}
		log15.Debug("Boot.Database.Close")
		return nil
	})
	return nil

}

func (b *Boot) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (b *Boot) Disable(ctx *pucore.ModuleContext) error {
	return pucore.ErrModuleDisableIgnore
}
