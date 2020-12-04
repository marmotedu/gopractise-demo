package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.V(3).Info("LEVEL 3 message") // 使用日志级别 3
	glog.V(5).Info("LEVEL 5 message") // 使用日志级别 5
	glog.V(7).Info("LEVEL 7 message") // 使用日志级别 7
	glog.V(8).Info("LEVEL 8 message") // 使用日志级别 8
}
