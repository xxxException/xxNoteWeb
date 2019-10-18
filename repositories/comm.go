package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"xxNoteWeb/readConf"
)

var db *sql.DB

func DbInit() {
	//sqlConf := "web/static"
	//scMap, err := readConf.ReadConf(sqlConf)
	//if nil != err {
	//	log.Fatal("read config fail ", err)
	//}
	var err error
	db, err = sql.Open("mysql", "root"+":"+"199762"+"@tcp"+"("+"localhost:3306"+")"+"/"+"xxNoteWeb"+"?"+"useSSL=false&serverTimezone=GMT%2B8")
	if nil != err {
		log.Fatal("fail init database errorDefine", err)
	}
	if db == nil {
		log.Fatal("fail init database", err)
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
	if db == nil {
		log.Fatal("db is nil")
	}
	return db
}
