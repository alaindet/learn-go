package gomitolo

import (
	"fmt"

	"github.com/joho/godotenv"
)

const version = "0.1.0"

type Gomitolo struct {
	AppName string
	Debug   bool
	Version string
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

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env", rootPath))
	if err != nil {
		return err
	}

	return nil
}

func (g *Gomitolo) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		dirPath := fmt.Sprintf("%s/%s", root, path)
		err := g.CreateDirIfNotExists(dirPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gomitolo) checkDotEnv(path string) error {
	filePath := fmt.Sprintf("%s/.env", path)
	err := g.CreateFileIfNotExists(filePath)
	if err != nil {
		return err
	}
	return nil
}
