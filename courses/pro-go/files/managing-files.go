package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func createDirectoriesExample() {
	here, err := os.Getwd()
	_ = err

	myPath := filepath.Join(here, "foo", "bar", "baz", "qez.json")

	// This creates all non-existing directories in the given path
	err = os.MkdirAll(filepath.Dir(myPath), 0766)
	_ = err

	file, err := os.OpenFile(myPath, os.O_CREATE|os.O_WRONLY, 0666)
	_ = err

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode([]string{"foo", "bar", "baz"})

	fmt.Printf("File written at %s\n", myPath)
}

func readDirectoryExample() {
	here, err := os.Getwd()
	_ = err

	dirEntries, err := os.ReadDir(here)
	_ = err

	for _, dirEntry := range dirEntries {
		fmt.Printf("Entry name: %v, IsDir: %v\n", dirEntry.Name(), dirEntry.IsDir())
	}
}

func checkFilesExistenceExample() {
	filesToCheck := []string{"nope.txt", "config.json"}

	for _, fileName := range filesToCheck {
		info, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Printf("ERROR: File %s does not exist\n", fileName)
			continue
		}
		if err != nil {
			fmt.Println("ERROR", err.Error())
			continue
		}
		fmt.Printf("Filename: %s, Size: %v", info.Name(), info.Size())
	}
}

func analyzeFilesExample() {
	here, err := os.Getwd()

	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}

	err = filepath.WalkDir(
		here,
		func(path string, dirEntry os.DirEntry, dirErr error) (err error) {
			info, _ := dirEntry.Info()
			fmt.Printf("Path %v\n", path)
			fmt.Printf("Size: %v bytes\n\n", info.Size())
			return
		},
	)
	_ = err
}

func managingFilesExamples() {
	// createDirectoriesExample()
	// readDirectoryExample()
	// checkFilesExistenceExample()
	analyzeFilesExample()
}
