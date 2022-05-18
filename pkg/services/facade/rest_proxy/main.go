package main

import (
	"flag"

	"github.com/golang/glog"

	proxy "github.com/jtomasevic/gRPCExample/pkg/services/facade/rest_proxy/proxy"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := proxy.Run(); err != nil {
		glog.Fatal(err)
	}
}
