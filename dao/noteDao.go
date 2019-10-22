package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xxNoteWeb/dataSource"
	"xxNoteWeb/datamodles"
	"xxNoteWeb/errorDefine"
)

type InoteDaoositories interface {
	QueryNote(symbol string) (*datamodles.Note, error)
	UpdateNote(symbol string, editTime string, content string) error
	InsertNote(symbol string, content string) error
}

type NoteDao struct {
	EngineGroup *xorm.EngineGroup
}

//config
func NewNoteDao() *NoteDao {
	return &NoteDao{
		EngineGroup: dataSource.GetEngineGroup(),
	}
}

func (noteDao *NoteDao) ExistNote(symbol string) (bool, error) {
	//看不懂，，
	//var re, err = noteDao.EngineGroup.Exist(&RecordExist{
	//	Name: "test1",
	//})
	has, err := noteDao.EngineGroup.SQL("select * from note where symbol = ?", symbol).Exist()
	return has, err
}

func (noteDao *NoteDao) QueryNote(symbol string) (*datamodles.Note, error) {
	var note *datamodles.Note

	has, err := noteDao.EngineGroup.Where("symbol = ?", symbol).Get(&note)
	if err != nil {
		return nil, err
	}

	if has == false {
		return nil, err
	}

	return note, err
}

func (noteDao *NoteDao) UpdateNote(symbol string, content string, editTime string) (err error) {
	re, err := noteDao.EngineGroup.Exec("update note set content = ?, editTime =? where symbol = ? ",
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

func (noteDao *NoteDao) InsertNote(symbol string, createTime string) error {
	var note = &datamodles.Note{Symbol: symbol, CreateTime: createTime}
	_, err := noteDao.EngineGroup.Insert(note)
	return err
}

func (noteDao *NoteDao) DeleteNode(symbol string) error {
	_, err := noteDao.EngineGroup.Exec("delete from note where symbol = ? ", symbol)
	if err != nil {
		return err
	}
	return err
}
