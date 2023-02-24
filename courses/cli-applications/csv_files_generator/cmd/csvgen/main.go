package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type config struct {
	lines int
	files int
	dir   string
}

// csvgen -lines 1000 -files 1000 -dir ./testdata
func main() {
	cfg := parseInput()
	err := run(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cfg config) error {
	dirPath, err := parseDir(cfg.dir)
	if err != nil {
		return err
	}

	// Create folders, if needed
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	for i := 0; i < cfg.files; i++ {
		filename := fmt.Sprintf("file%d.csv", i)
		err = createCsvFile(filepath.Join(dirPath, filename), cfg.lines)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseInput() config {

	lines := flag.Int("lines", 10, "How many lines a single .csv file must have")
	files := flag.Int("files", 10, "How many .csv files to generate")
	dir := flag.String("dir", ".", "Where to generate files")
	flag.Parse()

	return config{
		lines: *lines,
		files: *files,
		dir:   *dir,
	}
}

// Checks given directory path
// If it exists and it's a file, return error
// If it's an existing directory or a non-existing path, it's ok
// Returns absolute path if no error occurred
func parseDir(dir string) (string, error) {

	p, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	ps, err := os.Stat(p)

	// Non-existing path, it's ok
	if os.IsNotExist(err) {
		return p, nil
	}

	if err != nil {
		return "", err
	}

	mode := ps.Mode()

	// Maybe it's a file?
	if !mode.IsDir() {
		return "", errors.New("invalid path")
	}

	// It's an existing folder, it's ok
	return p, nil
}
