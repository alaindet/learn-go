package main

import "greenlight/cmd/api/core"

func main() {
	cfg := core.NewConfig()
	app := core.NewApplication(cfg)
	defer app.Shutdown()
	router := Routes(app, app.UrlPrefix)
	app.StartNewServer(router)
}
