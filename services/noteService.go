package services

import (
	"database/sql"
	"errors"
	"time"
	"xxNoteWeb/dao"
	"xxNoteWeb/datamodles"
	"xxNoteWeb/errorDefine"
)

type INoteService interface {
	InsertNode(string, string) error
}

type NoteService struct {
	dao *dao.NoteDao
	a   int
}

func NewNoteService() *NoteService {
	return &NoteService{
		dao: dao.NewNoteDao(),
		a:   2,
	}
}

func (noteSer *NoteService) GetNoteBySymbol(symbol string) (*datamodles.Note, error) {
	note, err := noteSer.dao.QueryNote(symbol)
	if err != sql.ErrNoRows {
		return nil, errorDefine.NoSymbolNoteErr
	}

	return note, err
}

func (noteSer *NoteService) NewNote(symbol string) error {
	time := getTime()

	err := noteSer.dao.InsertNote(symbol, time)
	if err != nil {
		return errorDefine.InsertNoteErr
	}
	return nil
}

func (noteSer *NoteService) UpdateNote(symbol string, content string) error {
	time := getTime()

	err := noteSer.dao.UpdateNote(symbol, content, time)
	if err != nil {
		return errorDefine.UpdateNoteErr
	}
	return nil
}

func (noteSer *NoteService) IsExistNote(symbol string) (bool, error) {
	has, err := noteSer.dao.ExistNote(symbol)
	if err != nil {
		return false, err
	}

	return has, err
}

func (noteSer *NoteService) DeleteNote(symbol string) error {
	err := noteSer.dao.DeleteNode(symbol)
	if err != nil {
		return errors.New("delete note fail : " + err.Error())
	}
	return nil
}

func getTime() string {
	time := time.Now().Local().String()
	time = time[:19]
	return time
}
