package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
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
	_ = err
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
	_ = err
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

func htmlTemplateWithNamedTemplates() {
	htmlTemplates, err := template.ParseGlob("templates/*.html")
	_ = err

	// NOTE
	// "mainTemplate" is a named template inside the named.html file!
	htmlTemplate := htmlTemplates.Lookup("mainTemplate")
	err = htmlTemplate.Execute(os.Stdout, &Products)
	_ = err
	// <p>This is a custom template from some-templates.html</p>
	// <p>There are 8 products in the source data.</p>
	// <p>First product: {Kayak Watersports 279}</p>
	// <ul>
	//   <li>
	//     [Midrange Product]
	//     Name: Kayak, Category: Watersports, Price: $279.00
	//   </li>
	//   <li>
	//     [Cheap Product]
	//     Name: Lifejacket, Category: Watersports, Price: $49.95
	//   </li>
	//   <li>
	//     [Cheap Product]
	//     Name: Soccer Ball, Category: Soccer, Price: $19.50
	//   </li>
	//   <li>
	//   	 [Cheap Product]
	//   	 Name: Corner Flags, Category: Soccer, Price:$34.95
	//   </li>
	//   <li>
	//   	 [Expensive Product]
	//   	 Name: Stadium, Category: Soccer, Price: $79500.00
	//   </li>
	//   <li>
	//   	 [Cheap Product]
	//   	 Name: Thinking Cap, Category: Chess, Price: $16.00
	//   </li>
	//   <li>
	//   	 [Cheap Product]
	//   	 Name: Unsteady Chair, Category: Chess, Price: $75.00
	//   </li>
	//   <li>
	//   	 [Midrange Product]
	//   	 Name: Bling-Bling King, Category: Chess, Price: $1200.00
	//   </li>
	// </ul>
}

func htmlTemplateWithBlocks() {
	// NOTE
	// Order is important since "blocks" can be overridden by "define"
	// (Ex.: in head.html) only if they have been already declared via "block"
	// (Ex.: in index.html)
	htmlTemplates, err := template.ParseFiles(
		"templates/blocks/index.html",
		"templates/blocks/head.html",
		"templates/blocks/body.html",
	)
	_ = err
	htmlTemplate := htmlTemplates.Lookup("mainTemplate")
	err = htmlTemplate.Execute(os.Stdout, &Products)
	_ = err
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	//   <meta charset="UTF-8">
	//   <meta http-equiv="X-UA-Compatible" content="IE=edge">
	//   <meta name="viewport" content="width=device-width, initial-scale=1.0">
	//   <title>Document</title>
	// </head>
	// <body>
	//   <p>This is the body</p>
	// </body>
	// </html>
}

func htmlTemplateWithCustomFunctions() {

	// This safely HTML-escaped
	// getCategories := func(products []Product) []string {
	// 	var categories []string
	// 	categoriesMap := map[string]string{}
	// 	for _, p := range products {
	// 		if categoriesMap[p.Category] == "" {
	// 			categoriesMap[p.Category] = p.Category
	// 			categories = append(categories, p.Category)
	// 		}
	// 	}
	// 	return categories
	// }

	// This is not HTML-escaped
	getCategories := func(products []Product) []template.HTML {
		var categories []template.HTML
		categoriesMap := map[string]string{}
		for _, p := range products {
			if categoriesMap[p.Category] == "" {
				categoriesMap[p.Category] = p.Category
				categories = append(categories, template.HTML(
					fmt.Sprintf("<strong>%s</strong>", p.Category),
				))
			}
		}
		return categories
	}

	templateFunctions := map[string]interface{}{
		"getCategories": getCategories,
		"lower":         strings.ToLower,
	}

	htmlTemplates := template.New("htmlTemplates")
	htmlTemplates.Funcs(templateFunctions)
	htmlTemplates, err := htmlTemplates.ParseGlob("templates/functions/*.html")
	_ = err

	htmlTemplate := htmlTemplates.Lookup("mainTemplate")
	err = htmlTemplate.Execute(os.Stdout, &Products)
	_ = err
	// <p>Some random content...</p>
	// <p>all caps! wait, what?</p>
	// <p>There are 3 categories</p>
	// <ul>
	//   <li>Category: <strong>Watersports</strong></li>
	//   <li>Category: <strong>Soccer</strong></li>
	//   <li>Category: <strong>Chess</strong></li>
	// </ul>
}

func htmlTemplatesExamples() {
	// singleHtmlTemplateExample()
	// multipleHtmlTemplatesExample()
	// listingHtmlTemplates()
	// lookUpHtmlTemplates()
	// htmlTemplatesWithLoops()
	// htmlTemplateWithNamedTemplates()
	// htmlTemplateWithBlocks()
	htmlTemplateWithCustomFunctions()
}
