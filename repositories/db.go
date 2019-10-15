package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"xxNoteWeb/services/readConf"
)

var db *sql.DB

func DbInit() {
	sqlConf := "/root/conf/sqlConf"
	scMap, err := readConf.ReadConf(sqlConf)

	db, err := sql.Open("mysql", scMap["username"]+":"+scMap["password"]+"@tcp"+"("+scMap["url"]+")"+"/"+scMap["dbname"]+"?"+scMap["ops"])
	if nil != err {
		panic("fail init database err")
	}
	if db == nil {
		panic("fail init database")
	}
}

func getDB() *sql.DB {
	return db
}
