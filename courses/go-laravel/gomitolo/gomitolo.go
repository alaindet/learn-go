package gomitolo

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alaindet/gomitolo/render"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "0.1.0"

type config struct {
	port     string
	renderer string
}

type Gomitolo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Render   *render.Render
	RootPath string
	Routes   *chi.Mux
	config   config
}

func (g *Gomitolo) New(rootPath string) error {

	err := g.Init(
		initPaths{
			rootPath: rootPath,
			folderNames: []string{
				"handlers",
				"migrations",
				"views",
				"data",
				"public",
				"tmp",
				"logs",
				"middleware",
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (g *Gomitolo) Init(p initPaths) error {

	g.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	g.Version = version
	g.RootPath = p.rootPath
	g.InitFolders(p)
	g.InitEnv()

	g.InfoLog, g.ErrorLog = g.createLoggers()
	g.config = g.createConfig()
	g.Render = g.createRenderer()
	g.Routes = g.routes().(*chi.Mux)

	return nil
}

func (g *Gomitolo) InitFolders(p initPaths) error {
	root := p.rootPath

	// Init folders
	for _, path := range p.folderNames {
		dirPath := fmt.Sprintf("%s/%s", root, path)
		err := CreateDirIfNotExists(dirPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Gomitolo) InitEnv() error {

	// .env.example
	envExamplePath := fmt.Sprintf("%s/.env.example", g.RootPath)
	if !FileExists(envExamplePath) {
		err := os.WriteFile(envExamplePath, []byte(envExampleTemplate), 0666)
		if err != nil {
			return err
		}
	}

	// .env
	envPath := fmt.Sprintf("%s/.env", g.RootPath)
	if !FileExists(envPath) {
		err := os.WriteFile(envPath, []byte(envExampleTemplate), 0666)
		if err != nil {
			return err
		}
	}

	// Load .env
	err := godotenv.Load(fmt.Sprintf("%s/.env", g.RootPath))
	if err != nil {
		return err
	}

	return nil
}

func (g *Gomitolo) createLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (g *Gomitolo) createConfig() config {
	return config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}
}

func (g *Gomitolo) createRenderer() *render.Render {
	return &render.Render{
		Renderer: g.config.renderer,
		RootPath: g.RootPath,
		Port:     g.config.port,
	}
}

func (g *Gomitolo) ListenAndServe() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     g.ErrorLog,
		Handler:      g.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	g.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))

	err := server.ListenAndServe()
	g.ErrorLog.Fatal(err)
}
