package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Context struct {
	Request *http.Request
	Data    []Product
}

var htmlTemplates *template.Template

func HandleTemplateRequest(w http.ResponseWriter, r *http.Request) {

	path := "products.html"

	if r.URL.Path != "" {
		path = r.URL.Path
	}

	htmlTemplate := htmlTemplates.Lookup(path)

	if htmlTemplate == nil {
		http.NotFound(w, r)
		return
	}

	err := htmlTemplate.Execute(w, Context{r, Products})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func initServer() {
	var err error
	htmlTemplates = template.New("all")

	/*
		Bind useful functions
	*/
	htmlTemplates.Funcs(map[string]interface{}{
		"intVal": strconv.Atoi,
	})

	htmlTemplates, err = htmlTemplates.ParseGlob("templates/*.html")

	if err != nil {
		panic(err)
	}

	httpHandler := http.StripPrefix(
		"/templates/",
		http.HandlerFunc(HandleTemplateRequest),
	)

	http.Handle("/templates/", httpHandler)
}

func dynamicTemplatesExample() {
	initServer()

	/*
		Bootstrap server
	*/
	fmt.Println("Starting HTTP server on :5000")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}
}
