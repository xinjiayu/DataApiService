package main

import (
	_ "DataApiService/boot"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
)

func main() {
	glog.Info("DataApiService Version:", "V0.0.1.2019071415")
	g.Server().Run()
}
