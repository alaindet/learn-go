package main

func main() {
	app := initApp()
	defer app.shutdown()

	app.infoLog.Printf("starting %s on %s\n", app.config.name, app.config.addr)
	server := app.initWebServer()
	err := server.ListenAndServe()
	app.errorLog.Fatal(err)
}
