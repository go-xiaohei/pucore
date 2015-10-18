package app

import (
	"github.com/Unknwon/com"
	"gopkg.in/ini.v1"
	"os"
)

const CONFIG_FILE = "config.ini"

type Config struct {
	Name    string
	Version string
	Date    string

	Http ConfigHttp
	Db   ConfigDb

	IsNew bool `ini:"-"`
}

type ConfigHttp struct {
	Host     string
	Port     string
	Protocol string
}

type ConfigDb struct {
	Driver string
	DSN    string
}

func NewConfig() *Config {
	config := &Config{
		Name:    "PUGO",
		Version: "2.0",
		Date:    "20151018",
		Http: ConfigHttp{
			Host:     "0.0.0.0",
			Port:     "9999",
			Protocol: "http",
		},
		Db: ConfigDb{
			Driver: "tidb",
			DSN:    "boltdb://data.db/tidb",
		},
	}
	return config
}

// Sync saves config to file.
// If config file is not exist, it creates new file and marks *Config is new file.
func (c *Config) Sync() error {
	// if config file exist, read it
	if com.IsFile(CONFIG_FILE) {
		file, err := ini.Load(CONFIG_FILE)
		if err != nil {
			return err
		}
		return file.MapTo(c)
	}
	// create file
	c.IsNew = true
	file := ini.Empty()
	if err := file.ReflectFrom(c); err != nil {
		return err
	}
	f, err := os.OpenFile(CONFIG_FILE, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = file.WriteToIndent(f, "  ")
	return err
}
