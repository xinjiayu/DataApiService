package table

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
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

var TableInfo g.MapStrStr

func GetTableInfo(tablename string) {

}

// GetAllTableInfo 获取所有表信息
func GetAllTableInfo(dbname string) {

	r, err := g.DB().Table("INFORMATION_SCHEMA.TABLES").Fields(
		"table_name as name,table_comment as comment").Where(
		"table_schema = ?", dbname).Select()
	if err != nil {
		glog.Error("gstart tables error", err)
	} else {
		TableInfo = g.MapStrStr{}
		list := r.ToList()
		for _, value := range list {
			TableInfo[gconv.String(value["name"])] = gconv.String(value["comment"])
		}
		glog.Info("gstart table info finish", TableInfo)
	}
}
