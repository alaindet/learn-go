package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func main() {

	// Input
	filename := flag.String("file", "", "Markdown file to process")
	showPreview := flag.Bool("preview", true, "Automatically open the HTML output file")
	userTemplate := flag.String("template", "", "Custom HTML template")
	flag.Parse()
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Execute
	err := run(*filename, *userTemplate, os.Stdout, *showPreview)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, userTemplate string, out io.Writer, showPreview bool) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(input, userTemplate)
	if err != nil {
		return err
	}

	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}

	defer os.Remove(temp.Name())
	err = temp.Close()
	if err != nil {
		return err
	}

	outputFilename := temp.Name()
	fmt.Fprintln(out, outputFilename)
	err = saveHTML(outputFilename, htmlData)
	if err != nil {
		return err
	}

	if !showPreview {
		return nil
	}

	return preview(outputFilename)
}

func parseContent(input []byte, userTemplate string) ([]byte, error) {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	// User-provided template?
	if userTemplate != "" {
		t, err = template.ParseFiles(userTemplate)
		if err != nil {
			return nil, err
		}
	}

	tParams := templateParams{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	var buffer bytes.Buffer
	err = t.Execute(&buffer, tParams)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}

func preview(filename string) error {
	cmdName := ""
	cmdParams := []string{}

	switch runtime.GOOS {
	case "linux":
		cmdName = "wslview"
	case "windows":
		cmdName = "cdm.exe"
		cmdParams = []string{"/C", "start"}
	case "darwin":
		cmdName = "open"
	default:
		return errors.New("OS not supported")
	}

	cmdParams = append(cmdParams, filename)
	cmdPath, err := exec.LookPath(cmdName)
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdPath, cmdParams...)
	err = cmd.Run()

	// This gives time to open the file
	time.Sleep(2 * time.Second)
	return err
}
