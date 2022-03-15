package main

import (
	"fmt"
	"os"
	"text/template"
)

func singleHtmlTemplateExample() {
	t, err := template.ParseFiles("templates/template.html")
	_ = err

	// Print parsed HTML template into standard output
	t.Execute(os.Stdout, &Kayak)
}

func multipleHtmlTemplatesExample() {
	htmlTemplatePaths := []string{
		"templates/template.html",
		"templates/extras.html",
	}

	htmlTemplates, err := template.ParseFiles(htmlTemplatePaths...)
	_ = err

	htmlTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
	os.Stdout.WriteString("\n\n======\n\n")
	htmlTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
}

func listingHtmlTemplates() {
	htmlTemplates, err := template.ParseGlob("templates/*.html")
	_ = err

	for _, htmlTemplate := range htmlTemplates.Templates() {
		fmt.Printf("Template name: %q\n", htmlTemplate.Name())
	}
}

func lookUpHtmlTemplates() {
	htmlTemplates, err := template.ParseGlob("templates/*.html")
	_ = err

	htmlTemplate := htmlTemplates.Lookup("extras.html")
	err = htmlTemplate.Execute(os.Stdout, &Kayak)
}

func htmlTemplatesExamples() {
	// singleHtmlTemplateExample()
	// multipleHtmlTemplatesExample()
	// listingHtmlTemplates()
	lookUpHtmlTemplates()
}
