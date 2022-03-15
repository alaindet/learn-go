package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func commonLocationFunctions() {

	// Current directory
	dir, err := os.Getwd()
	_ = err
	fmt.Println("os.Getwd()", dir)
	// os.Getwd() /home/alaindet/go/src/learn-go/courses/pro-go/files

	// User home directory
	userDir, err := os.UserHomeDir()
	_ = err
	fmt.Println("os.UserHomeDir()", userDir)
	// os.UserHomeDir() /home/alaindet

	// User cache directory
	cacheDir, err := os.UserCacheDir()
	_ = err
	fmt.Println("os.UserCacheDir()", cacheDir)
	// os.UserCacheDir() /home/alaindet/.cache

	// User config directory
	configDir, err := os.UserConfigDir()
	_ = err
	fmt.Println("os.UserConfigDir()", configDir)
	// os.UserConfigDir() /home/alaindet/.config

	// Default temporary dir
	tempDir := os.TempDir()
	_ = err
	fmt.Println("os.TempDir()", tempDir)
	// os.TempDir() /tmp
}

func manipulatePathsExamples() {
	// Append segments
	userDir, err := os.UserHomeDir()
	_ = err
	myPath := filepath.Join(userDir, "myapp", "myfile.json")
	fmt.Println("userdir + myapp + myfile.json =", myPath)

	// Absolute from relative
	absPath, err := filepath.Abs("hello.txt")
	_ = err
	fmt.Println("Absolute from relative 'hello.txt' =>", absPath)
	fmt.Println(filepath.IsAbs(absPath), filepath.IsAbs("hello.txt"))

	// Segments
	fmt.Println("Basename:", filepath.Base(absPath))
	fmt.Println("Directory:", filepath.Dir(absPath))
	fmt.Println("Extension:", filepath.Ext(absPath))

	// TODO: Change glob pattern?
	matches, err := filepath.Match("/**/hell*.txt", absPath)
	_ = err
	fmt.Println("Match?:", matches)

	// Split directory and filename
	dirSegment, fileSegment := filepath.Split(absPath)
	fmt.Printf("Dir: %s, File: %s\n", dirSegment, fileSegment)

	// Split all path in segments
	pathSegments := filepath.SplitList(absPath)
	fmt.Println(pathSegments)
}

func pathsExamples() {
	// commonLocationFunctions()
	manipulatePathsExamples()
}
