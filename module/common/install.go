package common

import (
	"fmt"
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
	"gopkg.in/inconshreveable/log15.v2"
)

type Install struct{}

func (is *Install) Name() string {
	return "INSTALL"
}

func (bs *Install) Bootstrap(inj *pucore.Injector) error {
	var (
		config   = new(app.Config)
		database = new(app.Db)
	)
	inj.Get(config)
	// if is new file, try to install
	if !config.IsNewFile {
		log15.Debug("AppService.Install.already")
		return nil
	}

	// set database
	inj.Get(database)
	if database.Engine == nil {
		if err := database.Connect(); err != nil {
			return err
		}
	}

	log15.Info("AppService.Install.Start")

	return nil
}

func (bs *Install) Depends(m *pucore.Modular) error {
	return nil
}

func (bs *Install) Enable() error {
	return nil
}

func (bs *Install) Disable() error {
	return nil
}
