package config

import (
	"log"

	"github.com/k-tsurumaki/todo-app/utils"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SLQDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("Port").MustString("8080"),
		SLQDriver: cfg.Section("db").Key("driver").String(),
		DBName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
