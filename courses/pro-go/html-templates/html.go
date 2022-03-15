package main

import (
	"fmt"
	"html/template"
	"os"
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

	htmlTemplate := htmlTemplates.Lookup("template.html")
	err = htmlTemplate.Execute(os.Stdout, &Kayak)
	// <ul>
	// 	<li>Template value as string: {Kayak Watersports 279}</li>
	// 	<li>Name: Kayak</li>
	// 	<li>Category: Watersports</li>
	// 	<li>Price: $279.00</li>
	// 	<li>Tax: $334.80</li>
	// 	<li>Discoun Price: $269.00</li>
	// </ul>
}

func htmlTemplatesWithLoops() {
	htmlTemplates, err := template.ParseGlob("templates/*.html")
	_ = err
	htmlTemplate := htmlTemplates.Lookup("loops.html")
	err = htmlTemplate.Execute(os.Stdout, &Products)
	// <p>There are 8 products</p>
	// <p>There are less than 10 products</p>
	// <p>First product: Kayak</p>
	// <ul>
	//   <li>Name: Kayak, Price: $279.00</li>
	//   <li>Name: Lifejacket, Price: $49.95</li>
	//   <li>Name: Soccer Ball, Price: $19.50</li>
	//   <li>Name: Corner Flags, Price: $34.95</li>
	//   <li>Name: Stadium, Price: $79500.00</li>
	//   <li>Name: Thinking Cap, Price: $16.00</li>
	//   <li>Name: Unsteady Chair, Price: $75.00</li>
	//   <li>Name: Bling-Bling King, Price: $1200.00</li>
	// </ul>

	// err = htmlTemplate.Execute(os.Stdout, &[]Product{})
	// <p>There are 0 products</p>
	// <p>There are less than 10 products</p>
	// <ul>
	//   <li>No products</li>
	// </ul>
}

func htmlTemplatesExamples() {
	// singleHtmlTemplateExample()
	// multipleHtmlTemplatesExample()
	// listingHtmlTemplates()
	// lookUpHtmlTemplates()
	htmlTemplatesWithLoops()
}
