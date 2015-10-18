package app

import (
	"github.com/go-xiaohei/pucore/core"
	_ "github.com/go-xorm/tidb"
	"github.com/ngaut/log"
	"github.com/pingcap/tidb"
	"gopkg.in/inconshreveable/log15.v2"
	"gopkg.in/inconshreveable/log15.v2/ext"
)

var (
	// global injector, saves all global variables such as config, database , server and so on,
	// it's used in module methods, should not be used in http context.
	Injector *pucore.Injector
	// global behavior, it maintains all behavior definition and hooks.
	Behavior *pucore.Behaviors
	// global modular manager. All modules should be registered, enabled or disable by the manager.
	Modular *pucore.Modular
)

func init() {
	// set log settings
	log15.Root().SetHandler(log15.LvlFilterHandler(log15.LvlDebug, ext.FatalHandler(log15.StderrHandler)))

	// close tidb debug
	log.SetLevelByString("error")
	tidb.Debug = false

	// init global vars
	var (
		config   *Config = NewConfig()
		database *Db     = NewDB(config.Db.Driver, config.Db.DSN)
	)
	Injector = pucore.NewInjector(config, database)
	Behavior = pucore.NewBehaviors()
	Modular = pucore.NewModular(Injector, Behavior)
}

type appService struct{}

func (as *appService) Start() {
	if err := Modular.EnableAll(); err != nil {
		log15.Crit("AppService.Start.Fail", "error", err)
	}
}

func (as *appService) Stop() {
	if err := Modular.DisableAll(); err != nil {
		log15.Crit("AppService.Stop.Fail", "error", err)
	}
}

func Run() {
	Start(new(appService))
}
