package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xxNoteWeb/dataSource"
	"xxNoteWeb/datamodles"
	"xxNoteWeb/errorDefine"
)

type INoteRepositories interface {
	QueryNote(symbol string) (*datamodles.Note, error)
	UpdateNote(symbol string, editTime string, content string) error
	InsertNote(symbol string, content string) error
}

type NoteRepository struct {
	EngineGroup *xorm.EngineGroup
}

//config
func NewNoteRepository() *NoteRepository {
	return &NoteRepository{
		EngineGroup: dataSource.NewMysqlEngineGroup(),
	}
}

func (noteRep *NoteRepository) ExistNote(symbol string) (bool, error) {
	//看不懂，，
	//var re, err = noteRep.EngineGroup.Exist(&RecordExist{
	//	Name: "test1",
	//})
	has, err := noteRep.EngineGroup.SQL("select symbol from note where symbol = ?", symbol).Exist()
	return has, err
}

func (noteRep *NoteRepository) QueryNote(symbol string) (*datamodles.Note, error) {
	var note *datamodles.Note

	has, err := noteRep.EngineGroup.Where("symbol = ?", symbol).Get(&note)
	if err != nil {
		return nil, err
	}

	if has == false {
		return nil, err
	}

	return note, err
}

func (noteRep *NoteRepository) UpdateNote(symbol string, content string, editTime string) (err error) {
	re, err := noteRep.EngineGroup.Exec("update note set content = ?, editTime =? where symbol = ? ",
		content, editTime, symbol)
	if err != nil {
		return err
	}

	//返回受影响的行数
	aff, _ := re.RowsAffected()
	if aff != 1 {
		//todo:rollback
		return errorDefine.RowsAffectMismatch
	}
	return nil
}

func (noteRep *NoteRepository) InsertNote(symbol string, createTime string) error {
	var note = &datamodles.Note{Symbol: symbol, CreateTime: createTime}
	_, err := noteRep.EngineGroup.Insert(note)
	return err
}
