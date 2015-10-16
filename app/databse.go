package app

import (
	_ "github.com/go-xorm/tidb"
	"github.com/go-xorm/xorm"
	"github.com/ngaut/log"
	"github.com/pingcap/tidb"
	"gopkg.in/inconshreveable/log15.v2"
)

func init() {
	log.SetLevelByString("error")
	tidb.Debug = false
}

type Db struct {
	*xorm.Engine

	Driver string
	DSN    string

	isConnected bool // if install process, it connects to database at first, so it may cause connect twice.
}

func (db *Db) Connect() error {
	if db.isConnected {
		return nil
	}
	engine, err := xorm.NewEngine(db.Driver, db.DSN)
	if err != nil {
		return err
	}
	if err := engine.Ping(); err != nil {
		return err
	}
	db.Engine = engine
	db.isConnected = true
	log15.Info("Database.connect." + db.Driver)
	return nil
}
