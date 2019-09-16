package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// Application initialization.
func init() {
	initConfig()
	initRouter()

}

// Configuration initialization.
func initConfig() {
	glog.Info("DataApiService start...")

	c := g.Config()
	s := g.Server()

	// log path
	logpath := c.GetString("setting.logPath")
	glog.SetPath(logpath)
	glog.SetStdoutPrint(true)

	// web Server configuration
	s.BindHookHandlerByMap("/*", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			r.Response.Header().Set("Access-Control-Allow-Origin", "*")
		},
	})

	//s.SetServerRoot("public")
	s.SetLogPath(logpath)
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(c.GetInt("setting.port"))
}
