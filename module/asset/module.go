package asset

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
	"github.com/lunny/tango"
)

const (
	STATIC_PREFIX string = "/static"
	STATIC_DIR    string = "static"
)

type Module struct {
}

func (tm *Module) Id() string {
	return "ASSET"
}

func (tm *Module) Prepare(ctx *pucore.ModuleContext) error {
	return nil
}

func (tm *Module) Enable(ctx *pucore.ModuleContext) error {
	static := &Static{
		Options: tango.StaticOptions{
			RootPath: STATIC_DIR,
			Prefix:   STATIC_PREFIX,
		},
	}

	var (
		server = new(app.Web)
	)
	ctx.Injector.Get(server)
	server.Use(static)
	ctx.Injector.Set(server)
	return nil
}

func (tm *Module) Disable(ctx *pucore.ModuleContext) error {
	return pucore.ErrModuleDisableIgnore
}
