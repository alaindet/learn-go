package main

type config struct {
	port string
}

type app struct {
	config *config
}

func NewApp() *app {
	return initApp()
}

func initConfig() *config {
	return &config{
		port: "4000", // TODO: Move to .env
	}
}

func initApp() *app {
	return &app{
		config: initConfig(),
	}
}
