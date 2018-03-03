package main

import (
	"flag"
	"github.com/ddliu/webhook/app"
	_ "github.com/ddliu/webhook/plugin"
)

func main() {
	var configFile = flag.String("config", "config.json", "Config file path")
	flag.Parse()

	app.NewApp(*configFile).Start()
}
