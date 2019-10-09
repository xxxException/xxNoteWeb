package datamodles

type Note struct {
	Id         int
	Symbol     string
	Content    string
	CreateTime string
	EditTime   string
}

func (note *Note) String() string {
	return "id: " + strconv.Itoa(note.Id) + " symbol: " + note.Symbol + " content: " + note.Content +
		" createTime: " + note.CreateTime + " editTime: " + note.EditTime
}
