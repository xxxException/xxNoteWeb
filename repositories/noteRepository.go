package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"noteWeb/datamodles"
	"noteWeb/services/readConf"
	"strconv"
)

type NodeRepositories interface {
	GetNote(symbol string) (datamodles.Note, error)
	UpdateNote(symbol string, editTime string, content string) error
}

func (db *sql.DB) GetNote(symbol string) (datamodles.Note, error) {
	if db == nil {
		panic("database do not init")
	}

	//begin tx
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sqlStr := "select id, symbol, content, createTime, editTime from note where symbol = ?"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(symbol)
	var id int
	var content string
	var symbol string
	var createTime string
	var editTime string
	err := rows.Scan(&id, &content, &symbol, &createTime, &editTime)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	note := datamodles.Note{Id: id, Content: content, Symbol: symbol, CreateTime: createTime, EditTime: editTime}
	tx.Commit()
	return note, nil
}

func (db *sql.DB) UpdateNote(symbol string, editTime string, content string) (err error) {
	if db == nil {
		panic("database do not init")
	}

	//begin tx
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sqlStr := "update note set content = ?, editTime =? where symbol = ?"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(content, editTime, symbol)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
