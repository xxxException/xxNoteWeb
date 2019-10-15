package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"xxNoteWeb/datamodles"
)

type INoteRepositories interface {
	QuertNote(symbol string) (*datamodles.Note, error)
	UpdateNote(symbol string, editTime string, content string) error
	InsertNote(symbol string, content string) error
}

type NoteRepository struct {
	DB *sql.DB
}

//config
func NewNoteRepository() *NoteReopsitory {
	return &NoteRepository{
		DB: getDB(),
	}
}

func (noteRep *NoteRepository) QueryNote(symbol string) (*datamodles.Note, error) {
	if noteRep.DB == nil {
		panic("database do not init")
	}

	//begin tx
	tx, err := noteRep.DB.Begin()
	if err != nil {
		return nil, err
	}

	sqlStr := "select id, symbol, content, createTime, editTime from note where symbol = ?"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(symbol)
	var id int
	var content string
	var createTime string
	var editTime string
	err = rows.Scan(&id, &content, &symbol, &createTime, &editTime)
	if err != nil { //sql.ErrNoRows
		tx.Rollback()
		return nil, err
	}
	note := datamodles.Note{Id: id, Content: content, Symbol: symbol, CreateTime: createTime, EditTime: editTime}
	tx.Commit()
	return &note, nil
}

func (noteRep *NoteRepository) UpdateNote(symbol string, editTime string, content string) (err error) {
	if noteRep.DB == nil {
		panic("database do not init")
	}

	//begin tx
	tx, err := noteRep.DB.Begin()
	if err != nil {
		return err
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
	if result != nil {
		return errors.New("update note fail :: symbol: " + symbol + " content: " + content)
	}
	tx.Commit()

	return nil
}

func (noteRep *NoteRepository) InsertNote(symbol string, content string, createTime string) error {
	if noteRep.DB == nil {
		panic("database do not init")
	}

	//begin tx
	tx, err := noteRep.DB.Begin()
	if err != nil {
		return err
	}

	sqlStr := "insert into note (symbol, content, createTime) values(?, ?, ?)"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(symbol, content, createTime)
	if err != nil {
		tx.Rollback()
		return err
	}
	if result != nil {
		return errors.New("insert note fail :: symbol: " + symbol + " content: " + content + " createTime:" + createTime)
	}
	tx.Commit()

	tx.Commit()

	return nil

}
