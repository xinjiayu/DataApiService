// Package api Data API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package table

import (
	"DataApiService/library/response"
	s_table "DataApiService/module/service/table"

	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

// Controller 控制器
type Controller struct {
}

// Get 获取指定表的数据。如果指定ID，将获取单条数据
// 参数说明：
// page：返回第几页
// limit：每一页多少条数据
// 默认设置为第1页，每页100条数据。
// field: 搜索哪儿个字段
// keyword: 搜索的值
func (ac *Controller) Get(r *ghttp.Request) {

	tableName := r.GetString("tablename")
	id := r.GetString("id")
	if id != "" {

		data, err1 := s_table.GetOneData(tableName, id)
		if err1 != nil {
			response.Json(r, 1, err1.Error())
		}
		response.Json(r, 0, "数据表："+tableName, data.ToMap())

	}

	page := r.GetInt("page")
	limit := r.GetInt("limit")
	field := r.GetString("field")
	keyword := r.GetString("keyword")
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	dataList, err := s_table.GetDataList(tableName, field, keyword, page, limit)
	if err != nil {
		response.Json(r, 1, err.Error())
	}

	response.Json(r, 0, "数据表："+tableName, dataList.ToList())

}

// Post 提交数据
func (ac *Controller) Post(r *ghttp.Request) {
	response.Json(r, 0, "RESTFul HTTP Method POST")

}

// Delete 删除数据
func (ac *Controller) Delete(r *ghttp.Request) {
	response.Json(r, 0, "RESTFul HTTP Method Delete")
}

// Put 更新指定的一条记录（提供该记录的全部信息）
func (ac *Controller) Put(r *ghttp.Request) {
	response.Json(r, 0, "RESTFul HTTP Method PUT")
}

// Patch 更新指定的一条记录（提供该记录的部分信息）
func (ac *Controller) Patch(r *ghttp.Request) {
	response.Json(r, 0, "RESTFul HTTP Method PATCH")
}

func (ac *Controller) Info(r *ghttp.Request) {
	tableName := r.GetString("tablename")
	s_table.GetTableInfo(tableName)

}

// Handler 处理器
func (ac *Controller) Handler(r *ghttp.Request) {
	glog.Info("========= Handler =======")
}

// HookHandler 处理器勾子
func (ac *Controller) HookHandler(r *ghttp.Request) {

	tableName := r.GetString("tablename")

	glog.Info("Hook Handler =======" + tableName)

}
