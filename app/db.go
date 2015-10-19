package app

import (
	"github.com/go-xorm/xorm"
)

type Db struct {
	*xorm.Engine
	DSN    string
	Driver string
}

func NewDB(driver, dsn string) *Db {
	return &Db{
		DSN:    dsn,
		Driver: driver,
	}
}

// Connect init *xorm.Engine really.
func (db *Db) Connect() error {
	engine, err := xorm.NewEngine(db.Driver, db.DSN)
	if err != nil {
		return err
	}
	if err := engine.Ping(); err != nil {
		return err
	}
	db.Engine = engine
	return nil
}
