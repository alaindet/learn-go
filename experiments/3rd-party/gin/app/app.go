package app

type Config struct {
	Port string
}

type App struct {
	Config
}

func NewApp() *App {
	app := &App{}
	app.init()
	return app
}

func (a *App) init() {
	a.initConfig()
}

func (a *App) initConfig() {
	a.Config = Config{
		Port: "4000",
	}
}
