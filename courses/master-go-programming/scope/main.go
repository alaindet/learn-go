package main

import (
	fmtAlias "fmt"
)

// Package-scoped
const done = false

func main() {

	// Block-scoped
	var taskIsRunning = false
	var done = true // This takes precedence over the package-scoped "done" constant

	fmtAlias.Println(taskIsRunning, done) // false true
	checkDoneStatus()
}

func checkDoneStatus() {
	// This uses the "done" package-scoped constant
	fmtAlias.Println(done) // false
}
