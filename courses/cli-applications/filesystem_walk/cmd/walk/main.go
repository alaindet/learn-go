package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	root    string
	ext     string
	size    int64
	list    bool
	del     bool
	logFile io.Writer
}

func main() {
	cfg := parseInput()

	err := run(cfg.root, os.Stdout, cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseInput() config {

	// Input
	root := flag.String("root", ".", "Root directory to start")
	logFile := flag.String("log", "", "Log deletes to this file")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	del := flag.Bool("del", false, "Delete files")
	flag.Parse()

	// Parse input
	f := os.Stdout

	if *logFile != "" {
		// os.O_APPEND: Enable appending data to existing file
		// os.O_CREATE: Create file if it doesn't exist
		// os.O_RDWR: Grants read and write permission
		f, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	return config{
		root:    *root,
		ext:     *ext,
		size:    *size,
		list:    *list,
		del:     *del,
		logFile: f,
	}
}

func run(rootDir string, out io.Writer, cfg config) error {

	deleteLogger := log.New(cfg.logFile, "DELETED FILE: ", log.LstdFlags)

	return filepath.Walk(
		rootDir,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if filterOutFile(path, cfg.ext, cfg.size, info) {
				return nil
			}

			if cfg.list {
				return listFile(path, out)
			}

			if cfg.del {
				return deleteFile(path, deleteLogger)
			}

			return listFile(path, out)
		})
}
