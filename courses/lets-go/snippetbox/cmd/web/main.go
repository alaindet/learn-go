package main

func main() {
	app := initApp()
	server := app.initWebServer()
	app.infoLog.Printf("starting %s on %s\n", app.config.name, app.config.addr)
	err := server.ListenAndServe()
	app.errorLog.Fatal(err)
}
