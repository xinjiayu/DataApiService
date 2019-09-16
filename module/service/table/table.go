package table

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var (
	// 数据库对象
	db = g.DB()
)

// GetOneData 获取一条数据记录
func GetOneData(tablename, id string) (gdb.Record, error) {

	table := db.Table(tablename).Safe()

	ret, err := table.Where("id=?", id).One()
	if err != nil {
		glog.Error(err)

	}
	return ret, err

}

// GetDataList 获取指定表的全部记录
func GetDataList(tablename, field, keyword string, page, limit int) (gdb.Result, error) {

	table := db.Table(tablename).Safe()
	var ret gdb.Result
	var err error

	if field != "" {
		ret, err = table.Where(field, keyword).ForPage(page, limit).Select()
	} else {
		ret, err = table.ForPage(page, limit).Select()
	}
	//ret,err := table.All()
	if err != nil {
		glog.Error(err)

	}
	return ret, err

}

func GetTableCount(tablename string) (int, error) {
	return g.DB().Table(tablename).Count()

}
