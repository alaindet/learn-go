package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := run(*filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	htmlData := parseContent(input)

	// The output file is in the same folder and has the same filename
	// it just has a .html suffix
	outputFilename := filename + ".html"

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
