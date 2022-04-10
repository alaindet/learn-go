package gomitolo

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

	return nil
}
