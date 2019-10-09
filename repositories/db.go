package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"noteWeb/services/readConf"
)

func DbInit() *sql.DB {
	sqlConf := "/root/conf/sqlConf"
	scMap, err := readConf.ReadConf(sqlConf)

	db, err := sql.Open("mysql", scMap["username"]+":"+scMap["password"]+"@tcp"+"("+scMap["url"]+")"+"/"+scMap["dbname"]+"?"+scMap["ops"])
	checkErr(err)
	return db
}
