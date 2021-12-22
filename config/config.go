package config

import (
	"gopkg.in/ini.v1"
)

type DbConf struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	DbConf DbConf
}

var GlobalConf *Config

func Load(confPath string) {
	cfg, err := ini.Load(confPath)
	if err != nil {
		panic(err)
	}

	GlobalConf = &Config{}
	GlobalConf.DbConf.Host = cfg.Section("db").Key("host").String()
	GlobalConf.DbConf.Port = cfg.Section("db").Key("port").String()
	GlobalConf.DbConf.User = cfg.Section("db").Key("user").String()
	GlobalConf.DbConf.Password = cfg.Section("db").Key("password").String()
	GlobalConf.DbConf.Database = cfg.Section("db").Key("database").String()
}
