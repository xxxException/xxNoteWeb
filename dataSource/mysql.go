package dataSource

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"xorm.io/xorm"
)

var EngineGroup *xorm.EngineGroup

func NewMysqlEngineGroup() *xorm.EngineGroup {
	var driverName = "mysql"
	var dataSourceName = "root:199762@tcp(127.0.0.1:199762)/xxNoteWeb?useSSL=false&serverTimezone=GMT%2B8"
	//one
	//engine, err := xorm.NewEngine(driverName, dataSourceName)
	//if err != nil {
	//	log.Fatal("xorm engine init fail: driver :", driverName, " dataSourceName : ", dataSourceName)
	//}

	//dataSourceNameSlice的数量，决定了EngineGroup内engine的数量
	var dataSourceNameSlice = []string{dataSourceName, dataSourceName, dataSourceName, dataSourceName, dataSourceName}
	EngineGroup, err := xorm.NewEngineGroup(driverName, dataSourceNameSlice)
	if err != nil {
		log.Fatal("xorm engine init fail: driver :", driverName, " dataSourceName : ", dataSourceName)
	}
	return EngineGroup
}
