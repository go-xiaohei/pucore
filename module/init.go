package module

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/module/asset"
	"github.com/go-xiaohei/pucore/module/auth"
	"github.com/go-xiaohei/pucore/module/boot"
	"github.com/go-xiaohei/pucore/module/setting"
	"github.com/go-xiaohei/pucore/module/theme"
	"github.com/go-xiaohei/pucore/module/user"
)

func init() {
	// register boot modules.
	// boot modules init basic variables to other modules, such as database, web server.
	// if install module finds install-process, call all boot.InstallModule to install.
	app.Modular.Register(new(boot.Boot), new(boot.Install), new(boot.Upgrade))

	// register base modules.
	// basic modules registers basic app variables to other modules, such as settings, theme, i18n
	app.Modular.Register(new(setting.Module), new(asset.Module), new(theme.Module))

	// register logic modules
	app.Modular.Register(new(user.Module), new(auth.Module))
}
