package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed web/dist
var web embed.FS

const defaultPort = "8080"

func main() {
	dist, _ := fs.Sub(web, "web/dist")
	http.Handle("/", http.FileServer(http.FS(dist)))

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(`{"data":"hello world"}`))
	})

	port := os.Getenv("port")
	if port == "" {
		port = defaultPort
	}

	address := fmt.Sprintf(":%s", port)
	fmt.Printf("Server started on port %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
