package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src := "source.txt"
	dest := "destination.txt"
	fmt.Printf("About to copy %q into %q...\n", src, dest)
	err := copyFile(src, dest)

	if err != nil {
		fmt.Println("ERROR", err)
	}
}

func copyFile(src, dest string) error {

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
