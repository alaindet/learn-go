package main

func main() {
	app := initApp()
	defer app.shutdown()

	app.infoLog.Printf("starting %s on %s\n", app.config.name, app.config.addr)
	server := app.initWebServer()
	err := server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	app.errorLog.Fatal(err)
}
