package common

import (
	"fmt"
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
)

type Bootstrap struct {
	server   *app.Server
	database *app.Db
}

func (bs *Bootstrap) Name() string {
	return "BOOTSTAP"
}

func (bs *Bootstrap) Bootstrap(inj *pucore.Injector) error {
	var (
		config   = new(app.Config)
		server   = new(app.Server)
		database = new(app.Db)
	)
	// read config
	inj.Get(config)

	// set server data
	inj.Get(server)
	server.Address = fmt.Sprintf("%s:%s", config.HttpHost, config.HttpPort)
	server.Domain = config.HttpDomain
	bs.server = server

	// set database
	inj.Get(database)
	database.Driver = config.DbDriver
	database.DSN = config.DbDSN
	bs.database = database

	inj.Set(server, database) // remember override value in injector

	return nil
}

func (bs *Bootstrap) Depends(m *pucore.Modular) error {
	return nil
}

func (bs *Bootstrap) Enable() error {
	bs.server.Start()
	if err := bs.database.Connect(); err != nil {
		return err
	}
	return nil
}

func (bs *Bootstrap) Disable() error {
	bs.server.Stop()
	bs.database.Clone()
	return nil
}
