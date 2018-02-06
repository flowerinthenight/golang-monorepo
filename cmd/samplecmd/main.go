package main

import (
	"flag"

	"github.com/golang/glog"
)

var version = "?"

func main() {
	flag.Parse()
	glog.CopyStandardLogTo("INFO")
	glog.Infof("Version: %v", version)
}
