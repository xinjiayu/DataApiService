package boot

import (
	a_table "DataApiService/module/api/table"
	a_user "DataApiService/module/api/user"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

/*
绑定业务路由
*/
func initRouter() {

	s := g.Server()

	// 用户模块 路由注册 - 使用执行对象注册方式
	s.BindObject("/v1/user/", new(a_user.Controller))

	obj := new(a_table.Controller)

	s.Group("/v1/data/:tablename").Bind([]ghttp.GroupItem{
		{"ALL", "*", obj.HookHandler, ghttp.HOOK_BEFORE_SERVE},
		{"ALL", "/handler", obj.Handler},
		{"GET", "/", obj, "Get"},
		{"GET", "{id}", obj, "Get"},
		{"DELETE", "{id}", obj, "Delete"},
		{"POST", "/", obj, "Post"},
		{"PUT", "{id}", obj, "Put"},
		{"PATCH", "{id}", obj, "Patch"},
		{"GET", "/count", obj, "Count"},
	})

}
