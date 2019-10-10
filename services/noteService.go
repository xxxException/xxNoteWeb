package services

import (
	"noteWeb/repositories"
	"time"
)

type INoteService interface {
	InsertNode(string, string) error
}

type NoteService struct {
	repositories.NoteRepositories
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
