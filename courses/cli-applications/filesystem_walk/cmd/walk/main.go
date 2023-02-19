package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	ext  string
	size int64
	list bool
}

func main() {

	// Input
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	cfg := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	err := run(*root, os.Stdout, cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(rootDir string, out io.Writer, cfg config) error {
	return filepath.Walk(
		rootDir,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			if cfg.list {
				return listFile(path, out)
			}

			return listFile(path, out)
		})
}
