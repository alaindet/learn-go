package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func filterOutFile(path, ext string, minSize int64, info os.FileInfo) bool {

	if info.IsDir() || info.Size() < minSize {
		return true
	}

	if ext != "" && filepath.Ext(path) != ext {
		return true
	}

	return false
}

func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}

func deleteFile(path string, deleteLogger *log.Logger) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	deleteLogger.Println(path)
	return nil
}
