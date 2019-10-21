package dataSource

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	////todo: go mod replace
	"xorm.io/xorm"
)

func NewMysqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine(
		conf.MysqlDriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			conf.MysqlMaster.User,
			conf.MysqlMaster.Pwd,
			conf.MysqlMaster.Host,
			conf.MysqlMaster.Port,
			conf.MysqlMaster.Database,
		),
	)
	return engine
}
