package main

import (
	"os"
)

func main() {
	config := readConfigFromFlags()
	app := initApplication(config)
	server := initServer(app)
	app.logger.Info("starting server", "addr", server.Addr, "env", config.env)
	err := server.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)
}
