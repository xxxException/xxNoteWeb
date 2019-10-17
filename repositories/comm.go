package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"xxNoteWeb/readConf"
)

var db *sql.DB

func DbInit() {
	sqlConf := "web/static/"
	scMap, err := readConf.ReadConf(sqlConf)

	db, err := sql.Open("mysql", scMap["username"]+":"+scMap["password"]+"@tcp"+"("+scMap["url"]+")"+"/"+scMap["dbname"]+"?"+scMap["ops"])
	if nil != err {
		panic("fail init database errorDefine")
	}
	if db == nil {
		panic("fail init database")
	}
}

func TxBegin(theDb *sql.DB) (*sql.Tx, error) {
	if theDb == nil {
		panic("database do not init")
	}

	//begin tx
	tx, err := theDb.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func getDB() *sql.DB {
	return db
}
