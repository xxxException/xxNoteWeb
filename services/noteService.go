package services

import (
	"database/sql"
	"noteWeb/repositories"
	"time"
)

type INoteService interface {
	InsertNode(string, string) error
}

type NoteService struct {
	repositories.NoteRepositories
}

func (noteSer *NoteService) GetNoteBySymbol(symbol string) {
	note, err := noteSer.QueryNote(symbol)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return note, nil
}

func (noteSer *NoteService) InsertNote(symbol string, content string) error {
	time := time.Now().Local().String()
	time = time[:19]

	err := noteSer.InsertNote(symbol, content)
	if err != nil {
		return err
	}
	return nil
}
