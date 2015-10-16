package module

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/module/common"
)

func init() {
	// add Bootstrap modular as default,
	// it bootstraps global vars with correct config data
	app.Modular.Register(new(common.Bootstrap))
	// app.Modular.Enable(new(common.Bootstrap).Name())

	// add Install modular,
	// try to begin install process if first run,
	// or skip it
	app.Modular.Register(new(common.Install))
}
