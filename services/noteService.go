package services

import (
	"database/sql"
	"time"
	"xxNoteWeb/dao"
	"xxNoteWeb/datamodles"
	"xxNoteWeb/errorDefine"
)

type INoteService interface {
	InsertNode(string, string) error
}

type NoteService struct {
	rep *dao.NoteRepository
}

func NewNoteService() *NoteService {
	return &NoteService{
		rep: dao.NewNoteRepository(),
	}
}

func (noteSer *NoteService) GetNoteBySymbol(symbol string) (*datamodles.Note, error) {
	note, err := noteSer.rep.QueryNote(symbol)
	if err != sql.ErrNoRows {
		return nil, errorDefine.NoSymbolNoteErr
	}

	return note, err
}

func (noteSer *NoteService) NewNote(symbol string) error {
	time := getTime()

	err := noteSer.rep.InsertNote(symbol, time)
	if err != nil {
		return errorDefine.InsertNoteErr
	}
	return nil
}

func (noteSer *NoteService) UpdateNote(symbol string, content string) error {
	time := getTime()

	err := noteSer.rep.UpdateNote(symbol, content, time)
	if err != nil {
		return errorDefine.UpdateNoteErr
	}
	return nil
}

func (noteSer *NoteService) IsExistNote(symbol string) (bool, error) {
	count, err := noteSer.rep.CountNote(symbol)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, err
	}
	return false, err
}

func getTime() string {
	time := time.Now().Local().String()
	time = time[:19]
	return time
}
