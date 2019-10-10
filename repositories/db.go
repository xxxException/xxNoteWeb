package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"noteWeb/services/readConf"
)

var DB *sql.DB

func DbInit() {
	sqlConf := "/root/conf/sqlConf"
	scMap, err := readConf.ReadConf(sqlConf)

	DB, err := sql.Open("mysql", scMap["username"]+":"+scMap["password"]+"@tcp"+"("+scMap["url"]+")"+"/"+scMap["dbname"]+"?"+scMap["ops"])
	if nil != err {
		panic("fail init database err")
	}
	if DB == nil {
		panic("fail init database")
	}
}
