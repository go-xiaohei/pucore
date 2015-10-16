package app

import (
	"github.com/go-xiaohei/pucore/core"
	"gopkg.in/inconshreveable/log15.v2"
	"gopkg.in/inconshreveable/log15.v2/ext"
)

var (
	Injector *pucore.Injector
	Modular  *pucore.Modular
)

func init() {
	// set log settings
	log15.Root().SetHandler(log15.LvlFilterHandler(log15.LvlDebug, ext.FatalHandler(log15.StderrHandler)))

	// make default global vars to injector
	// so all modules can use the global variables
	var (
		config   *Config = newConfig()
		server   *Server = newServer()
		database *Db     = new(Db)
	)
	Injector = pucore.NewInjector(config, server, database)
	Modular = pucore.NewModular(Injector)
}

// appService implements signal Service interface
type appService struct{}

func (a *appService) Start() {
	if err := Modular.Enable(); err != nil {
		log15.Crit("AppService.Start.Fail", "error", err)
	}
}

func (a *appService) Stop() {
	if err := Modular.Disable(); err != nil {
		log15.Crit("AppService.Stop.Fail", "error", err)
	}
}

func Run() {
	// run signal-sensitive service
	Start(new(appService))
}
