package main

import (
	"crypto/tls"
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
	sessionKeyFlash  = "flash"
	sessionKeyUserId = "userId"
)

type application struct {
	config
	errorLog       *log.Logger
	infoLog        *log.Logger
	db             *sql.DB // TODO: Remove?
	pgxPool        *pgxpool.Pool
	snippets       *models.SnippetModel
	users          *models.UserModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func initApp() *application {

	// Init loggers
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Init config
	config := NewConfig()

	// Init database
	db, err := openDB(config.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Init templates
	templateCache, err := newTemplateCache()
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
		users:          models.NewUserModel(db),
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
	}
}

func (app *application) initWebServer() *http.Server {

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.X25519,
			tls.CurveP256,
		},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	return &http.Server{
		Addr:         app.config.addr,
		ErrorLog:     app.errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (app *application) shutdown() {
	app.db.Close()
	app.pgxPool.Close()
}
