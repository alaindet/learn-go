package main

import (
	"fmt"
	"net/http"
	"strings"
)

const certFile = "certificates/certificate.cert"
const certKey = "certificates/certificate.key"
const httpPort = "5000"
const httpsPort = "5500"

func HTTPSServerExample() {
	/*
		Declare handlers
	*/
	http.Handle("/message", StringHandler{"Hello World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	/*
		Start HTTPS server
	*/
	go func() {
		fmt.Printf("Starting HTTPS server on :%s\n", httpsPort)
		addr := fmt.Sprintf(":%s", httpsPort)
		err := http.ListenAndServeTLS(addr, certFile, certKey, nil)

		if err != nil {
			fmt.Printf("HTTPS Error: %v\n", err.Error())
			return
		}
	}()

	/*
		Start HTTP server, redirect to HTTPS
	*/
	fmt.Printf("Starting HTTP server on :%s\n", httpPort)
	addr := fmt.Sprintf(":%s", httpPort)
	err := http.ListenAndServe(addr, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			query := ""
			if len(r.URL.RawQuery) > 0 {
				query = fmt.Sprintf("?%s", r.URL.RawQuery)
			}
			host := strings.Split(r.Host, ":")[0]
			url := fmt.Sprintf("https://%s:%s%s%s", host, httpsPort, r.URL.Path, query)

			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		},
	))

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}
}
