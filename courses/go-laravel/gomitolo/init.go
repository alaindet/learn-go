package gomitolo

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func (g *Gomitolo) Init(p initPaths) error {

	InitFolders(p)
	InitEnv(p)

	return nil
}

func (g *Gomitolo) InitEnv(p initPaths) error {
	root := p.rootPath

	// .env.example
	envExamplePath := fmt.Sprintf("%s/.env.example", root)
	if !FileExists(envExamplePath) {
		err := os.WriteFile(envExamplePath, []byte(envExampleTemplate), 0666)
		if err != nil {
			return err
		}
	}

	// .env
	envPath := fmt.Sprintf("%s/.env", path)
	err := CreateFileIfNotExists(envPath)
	if err != nil {
		return err
	}

	// Load
	err = godotenv.Load(fmt.Sprintf("%s/.env", root))
	if err != nil {
		return err
	}

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
