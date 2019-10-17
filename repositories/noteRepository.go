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
func NewNoteRepository() *NoteRepository {
	return &NoteRepository{
		DB: getDB(),
	}
}

func (noteRep *NoteRepository) CountNote(symbol string) (int, error) {
	tx, err := TxBegin(noteRep.DB)
	if err != nil {
		return -1, err
	}

	sqlStr := "select count(symbol) from note where symbol = ?"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(symbol)
	var count int
	err = rows.Scan(&count)
	if err != nil { //sql.ErrNoRows
		tx.Rollback()
		return -1, err
	}

	tx.Commit()
	return count, nil
}

func (noteRep *NoteRepository) QueryNote(symbol string) (*datamodles.Note, error) {
	tx, err := TxBegin(noteRep.DB)
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

func (noteRep *NoteRepository) UpdateNote(symbol string, content string, editTime string) (err error) {
	tx, err := TxBegin(noteRep.DB)
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

func (noteRep *NoteRepository) InsertNote(symbol string, createTime string) error {
	tx, err := TxBegin(noteRep.DB)
	if err != nil {
		return err
	}

	sqlStr := "insert into note (symbol, createTime) values(?, ?)"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(symbol, createTime)
	if err != nil {
		tx.Rollback()
		return err
	}
	if result != nil {
		return errors.New("insert note fail :: symbol: " + symbol + " createTime:" + createTime)
	}
	tx.Commit()

	tx.Commit()

	return nil

}
