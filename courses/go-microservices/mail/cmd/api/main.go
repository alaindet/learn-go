package main

import (
	"common/json"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

const (
	webPort = "80"
)

type App struct {
	json.HTTPClient
	Mailer Mail
}

func main() {
	log.Printf("Starting mail service on port %s\n", webPort)

	// TODO: Remove
	inspectEmbeddedFiles(embeddedTemplates)

	mailer, err := NewMail()
	if err != nil {
		log.Panic(err)
	}

	app := App{
		Mailer: mailer,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

// TODO: Remove
func inspectEmbeddedFiles(embeddedFs embed.FS) {
	fmt.Println("--- Inspecting Embedded Files ---")

	// fs.WalkDir lets you traverse the embedded file system
	err := fs.WalkDir(embeddedFs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Printf("[Directory] %s\n", path)
		} else {
			fmt.Printf("[File]      %s\n", path)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking embedded FS: %v", err)
	}
	fmt.Println("---------------------------------")
}
