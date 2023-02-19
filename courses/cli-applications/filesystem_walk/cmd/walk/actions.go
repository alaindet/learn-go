package main

import (
	"compress/gzip"
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

func archiveFile(destDir, root, path string) error {
	info, err := os.Stat(destDir)

	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", destDir)
	}

	relDir, err := filepath.Rel(root, filepath.Dir(path))

	if err != nil {
		return err
	}

	dest := fmt.Sprintf("%s.gz", filepath.Base(path))
	targetPath := filepath.Join(destDir, relDir, dest)
	err = os.MkdirAll(filepath.Dir(targetPath), 0755)

	if err != nil {
		return err
	}

	out, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer out.Close()
	in, err := os.Open(path)

	if err != nil {
		return err
	}

	defer in.Close()
	zw := gzip.NewWriter(out)
	zw.Name = filepath.Base(path)
	_, err = io.Copy(zw, in)

	if err != nil {
		return err
	}

	err = zw.Close()
	if err != nil {
		return err
	}

	return out.Close()
}
