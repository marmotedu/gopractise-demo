package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := loadConfig(); err != nil {
		glog.Error(err)
	}
}

func loadConfig() error {
	return decodeConfig()
}

func decodeConfig() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("could not decode configuration data for user %s: %v", "colin", err)
	}

	return nil
}

func readConfig() error {
	glog.Errorf("read: end of input.")
	return fmt.Errorf("read: end of input")
}
