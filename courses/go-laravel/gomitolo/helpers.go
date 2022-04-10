package gomitolo

import "os"

func (g *Gomitolo) CreateDirIfNotExists(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		const mode = 0755
		err := os.Mkdir(path, mode)

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Gomitolo) CreateFileIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)

		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}
