package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func main() {

	// Input
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Execute
	err := run(*filename, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}

	err = temp.Close()
	if err != nil {
		return err
	}

	outputFilename := temp.Name()
	fmt.Fprintln(out, outputFilename)
	return saveHTML(outputFilename, htmlData)
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var buffer bytes.Buffer
	buffer.WriteString(headerHTML)
	buffer.Write(body)
	buffer.WriteString(footerHTML)
	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}
