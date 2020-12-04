package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	glog.MaxSize = 1024 * 1024 * 1024 // 1G自动分割
	flag.Parse()
	defer glog.Flush()

	glog.Info("This is info message")
	glog.Infof("This is info message: %v", 123)

	glog.Warning("This is warning message")
	glog.Warningf("This is warning message: %v", 123)

	glog.Error("This is error message")
	glog.Errorf("This is error message: %v", 123)

	//glog.Fatal("This is fatal message")
	//glog.Fatalf("This is fatal message: %v", 123)
}
