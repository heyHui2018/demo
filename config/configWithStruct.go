package main

import (
	"github.com/go-ini/ini"
	"strings"
	"fmt"
	"os"
)

var ConfigInstance *Config

type FtpConfig struct {
	Host string
	Port string
}

type LogConfig struct {
	Path  string
	Level string
}

type Config struct {
	Ftp FtpConfig
	Log LogConfig
}

func main() {
	ConfigInit()
	fmt.Println("done-----------------------------------------")
}

func ConfigInit() {
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}
	config_file := "config.conf"
	config_file = strings.Replace(config_file, "/", separator, -1)
	c, err := ini.Load(config_file)
	if err != nil {
		panic(err)
	}

	ConfigInstance = new(Config)
	logConfig(c)
	ftpConfig(c)
	fmt.Printf("config = %+v", *ConfigInstance)
}

func logConfig(c *ini.File) {
	log, err := c.GetSection("log")
	handleErr(err)
	path, err := log.GetKey("path")
	handleErr(err)
	ConfigInstance.Log.Path = path.String()
	level, err := log.GetKey("level")
	handleErr(err)
	ConfigInstance.Log.Level = level.String()
}

func ftpConfig(c *ini.File) {
	ftp, err := c.GetSection("ftp")
	handleErr(err)
	host, err := ftp.GetKey("host")
	handleErr(err)
	ConfigInstance.Ftp.Host = host.String()
	port, err := ftp.GetKey("port")
	handleErr(err)
	ConfigInstance.Ftp.Port = port.String()
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
