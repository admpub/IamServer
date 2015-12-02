package main

import (
	"flag"
	"net/http"
	"os"
	"runtime/pprof"

	myhttp "github.com/Alienero/IamServer/http"
	"github.com/Alienero/IamServer/im"

	"github.com/golang/glog"
)

var isDebug = true

func main() {
	if isDebug {
		f, err := os.Create("pprof")
		if err != nil {
			glog.Fatal(err)
		}
		if err = pprof.StartCPUProfile(f); err != nil {
			glog.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}
	flag.Parse()
	defer glog.Flush()
	if err := flag.Set("logtostderr", "true"); err != nil {
		panic(err)
	}
	r := NewSrsServer()
	r.PrintInfo()
	// init http server
	if err := myhttp.InitHTTP(); err != nil {
		panic(err)
	}
	im.GlobalIM.Init()
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			panic(err)
		}
	}()
	r.Serve()
}
