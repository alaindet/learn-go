package gomitolo

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dest string) error {

	srcStat, err := os.Stat(src)

	if err != nil {
		return err
	}

	if !srcStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)

	if err != nil {
		return err
	}

	defer source.Close()

	destFile, err := os.Create(dest)

	if err != nil {
		return err
	}

	defer destFile.Close()

	nBytes, err := io.Copy(destFile, source)
	_ = nBytes

	return err
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
