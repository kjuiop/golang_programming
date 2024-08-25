package main

import (
	"chat_controller/cmd/app"
	"chat_controller/config"
	"flag"
)

var pathFlag = flag.String("config", "../config.toml", "config set")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*pathFlag)

	a := app.NewApp(cfg)
	a.Start()
}
