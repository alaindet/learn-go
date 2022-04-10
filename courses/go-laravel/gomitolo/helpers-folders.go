package gomitolo

import "os"

func CreateDirIfNotExists(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		const mode = 0755
		err := os.Mkdir(path, mode)

		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFileIfNotExists(path string) error {
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
