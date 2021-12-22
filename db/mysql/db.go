package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"stock/config"
)

var db *gorm.DB

func Init() {
	var (
		dbUser     = config.GlobalConf.DbConf.User
		dbPassword = config.GlobalConf.DbConf.Password
		dbHost     = config.GlobalConf.DbConf.Host
		dbPort     = config.GlobalConf.DbConf.Port
		dbName     = config.GlobalConf.DbConf.Database
	)
	var err error
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("connect db err:" + err.Error())
	}
}

func Migrate() {
	err := db.AutoMigrate(&Stock{})
	if err != nil {
		panic("auto migrate db err: " + err.Error())
	}
	// 创建表时添加后缀
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Stock{})
}
