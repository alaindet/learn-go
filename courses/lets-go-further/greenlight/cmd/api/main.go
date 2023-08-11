package main

func main() {
	cfg := NewConfig()
	app := NewApplication(cfg)
	defer app.Shutdown()
	app.StartNewServer()
}
