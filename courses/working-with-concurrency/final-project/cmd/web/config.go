package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"

	"final_project/data"
)

const (
	dbMaxTries   = 10
	dbRetryIn    = 1000 // milliseconds
	webPort      = "8080"
	mailerSecret = "abc123abc123abc123"
)

var db *sql.DB
var templatesPath = "./cmd/web/templates"
var tempPath = "./tmp"
var pdfTemplatesPath = "./pdf"

type Config struct {
	Session       *scs.SessionManager
	DB            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Wait          *sync.WaitGroup
	Models        data.Models
	Mailer        Mail
	ErrorChan     chan error
	ErrorChanDone chan bool
}
