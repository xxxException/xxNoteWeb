package datamodles

import (
	"strconv"
)

type Note struct {
	Id         int
	Symbol     string
	Content    string
	CreateTime string `xorm:"DATETIME 'create_time'"`
	EditTime   string `xorm:"DATETIME 'edit_time'"`
}

func (note *Note) String() string {
	return "id: " + strconv.Itoa(note.Id) + " symbol: " + note.Symbol + " content: " + note.Content +
		" createTime: " + note.CreateTime + " editTime: " + note.EditTime
}
