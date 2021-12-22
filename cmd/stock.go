package main

import (
	"flag"
	"stock/cmd/commad"
	"stock/config"
	"stock/db"
	"stock/db/mysql"
	"stock/router"
)

var (
	confPath = flag.String("conf", "/usr/local/stock/etc/stock.ini", "conf file")
	mode     = flag.Bool("mode", false, "api or command")
	migrate  = flag.Bool("migrate", false, "migrate db")
)

func Init() {
	flag.Parse()

	config.Load(*confPath)
	mysql.Init()
}

func main() {
	Init()

	// init srv
	dbSrv := db.NewMysqlService()

	if *mode {
		commad.Command(dbSrv)
		return
	}
	router.Server(dbSrv)
}
