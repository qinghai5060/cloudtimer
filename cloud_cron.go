package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"src/CloudCron/cmd"

	"github.com/golang/glog"
)

const VultrKey = "API_KEY"

var confFile string

func init() {
	flag.StringVar(&confFile, "conf", "conf.yaml", "path of config file")
}

func main() {
	flag.Parse()
	glog.CopyStandardLogTo("INFO")
	defer glog.Flush()

	glog.Info("start to work...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	cc := cmd.NewCloudCron(ctx, os.Getenv(VultrKey))
	if err := cc.Run(ctx, confFile); err != nil {
		panic(err)
	}

	<-ch
}
