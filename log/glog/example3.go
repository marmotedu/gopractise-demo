package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Info("This is info message")
	glog.Infof("This is info message: %v", 123)
}
