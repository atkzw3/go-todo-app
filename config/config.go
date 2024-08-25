package config

import (
	"gopkg.in/go-ini/ini.v1"
	"log"
	"todo-app/utils"
)

type List struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config List

func init() {
	// LoadConfig main関数より前に発火するように
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = List{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
