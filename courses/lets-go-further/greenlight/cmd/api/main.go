package main

func main() {
	cfg := NewConfig()
	app := NewApplication(cfg)
	server := NewServer(app.routes(), cfg)
	app.Start(server)
}
