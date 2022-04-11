package gomitolo

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "0.1.0"

type Gomitolo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
}

type config struct {
	port     string
	renderer string
}

func (g *Gomitolo) New(rootPath string) error {
	pathConfig := initPaths{
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
	}

	err := g.Init(pathConfig)

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
	g.InfoLog, g.ErrorLog = g.InitLoggers()

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

func (g *Gomitolo) InitLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
