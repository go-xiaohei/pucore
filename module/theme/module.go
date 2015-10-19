package theme

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
	"github.com/go-xiaohei/pucore/module/setting"
)

type Module struct{}

func (tm *Module) Id() string {
	return "THEME"
}

func (tm *Module) Prepare(ctx *pucore.ModuleContext) error {
	if !ctx.Modular.Has(new(setting.Module)) {
		return pucore.ErrModuleDependsOn(tm, new(setting.Module))
	}
	return nil
}

func (tm *Module) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (tm *Module) Disable(ctx *pucore.ModuleContext) error {
	return pucore.ErrModuleDisableIgnore
}

func (tm *Module) Install(ctx *pucore.ModuleContext) error {
	database := new(app.Db)
	ctx.Injector.Get(database)

	// insert theme setting as default setting
	setting := &setting.Setting{
		Name:   "theme",
		UserId: 9,
	}
	setting.Encode(&Theme{
		Name:        "pugo",
		Description: "pugo's default beautiful and responsive theme",
		Homepage:    "https://github.com/go-xiaohei/pugo",
		Tags:        []string{"blog", "rss", "syntax highlighting"},
		License:     "MIT",
		MinVersion:  "2.0",
		Author: struct {
			Name     string
			Homepage string
		}{
			Name:     "fuxiaohei",
			Homepage: "https://github.com/fuxiaohei",
		},
	})
	if _, err := database.Insert(setting); err != nil {
		return err
	}
	return nil
}
