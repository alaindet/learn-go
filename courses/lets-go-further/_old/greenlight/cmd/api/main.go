package main

func main() {
	cfg := NewConfig()
	app := NewApplication(cfg)
	defer app.Shutdown()
	server := NewServer(app.routes(), cfg)
	app.Start(server)
}
