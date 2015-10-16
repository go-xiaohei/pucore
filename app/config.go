package app

import (
	"github.com/Unknwon/com"
	"gopkg.in/ini.v1"
	"io"
	"os"
)

var (
	CONFIG_FILE = "config.ini"
)

type Config struct {
	Name    string
	Version string
	Date    string

	HttpProtocol string
	HttpHost     string
	HttpPort     string
	HttpDomain   string

	DbDriver string
	DbDSN    string

	IsNewFile bool `ini:"-"`
}

// new config data,
// if config file exist, read file;
// otherwise, create new file and mark IsNewFile flag
func newConfig() *Config {
	c := &Config{
		Name:    "PUGO",
		Version: "2.0",
		Date:    "20151016",

		HttpProtocol: "http",
		HttpHost:     "0.0.0.0",
		HttpPort:     "9999",
		HttpDomain:   "localhost",

		DbDriver: "tidb",
		DbDSN:    "boltdb://data.db/tidb",
	}
	// if config file is not exist,
	// set as new file and write to new file
	if !com.IsFile(CONFIG_FILE) {
		c.IsNewFile = true
		file, err := os.OpenFile(CONFIG_FILE, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if err := c.WriteTo(file); err != nil {
			panic(err)
		}
		return c
	}
	// read from file
	file, err := ini.Load(CONFIG_FILE)
	if err != nil {
		panic(err)
	}
	if err := file.MapTo(c); err != nil {
		panic(err)
	}
	return c
}

// write config ini bytes to io.Writer
func (c *Config) WriteTo(w io.Writer) error {
	file := ini.Empty()
	if err := file.ReflectFrom(c); err != nil {
		panic(err)
	}
	_, err := file.WriteToIndent(w, "  ")
	return err
}
