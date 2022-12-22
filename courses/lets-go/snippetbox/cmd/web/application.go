package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"snippetbox.dev/internal/models"
)

var (
	flashKey = "flash"
)

type application struct {
	config
	errorLog       *log.Logger
	infoLog        *log.Logger
	db             *sql.DB // TODO: Remove?
	pgxPool        *pgxpool.Pool
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func initApp() *application {

	// Init loggers
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Init config
	config := *NewConfig()

	// Init database
	db, err := openDB(config.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Init templates
	templateCache, err := newTemplateCache(config.htmlTemplatesPath)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Init session manager
	pgxDb, err := openPgxDB(config.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(pgxDb)
	sessionManager.IdleTimeout = 30 * time.Minute
	sessionManager.Lifetime = 3 * time.Hour

	// Application
	return &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		config:         config,
		db:             db,
		snippets:       models.NewSnippetModel(db),
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
	}
}

func (app *application) initWebServer() *http.Server {
	return &http.Server{
		Addr:     app.config.addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}
}

func (app *application) shutdown() {
	app.db.Close()
	app.pgxPool.Close()
}
