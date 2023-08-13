package main

import "greenlight/cmd/api/core"

func main() {
	cfg := core.NewConfig()
	app := core.NewApplication(cfg)
	defer app.Shutdown()
	prefix := "/" + core.ApiVersion
	router := Routes(app, prefix)
	app.StartNewServer(router)
}
