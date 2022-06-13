package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"

	"final_project/data"
)

var testApp Config

// https://pkg.go.dev/testing#hdr-Main
// https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
// - This will be called instead of running tests directly
// - Only 1 TestMain per package can exist
// - Use testing.M.Run() to run tests and get exit code
// - Remember to call os.Exit(exitCode) with exitCode != 0
func TestMain(m *testing.M) {
	initTestApp()
	go listenToTestMailer()
	go listenToTestLoggers()
	os.Exit(m.Run())
}

func initTestApp() {

	mailer := initTestMailer()
	infoLog, errorLog := initLoggers()

	testApp = Config{
		Session:       initTestSession(),
		DB:            nil,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Wait:          &sync.WaitGroup{},
		Models:        data.TestNew(nil),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		Mailer:        *mailer,
	}
}

func initTestSession() *scs.SessionManager {
	// Register a non-primitive Go data type so that it can be serialized
	// as GOB format into the session (Redis)
	gob.Register(data.User{})

	// Mock session
	session := scs.New()
	// session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initTestMailer() *Mail {
	return &Mail{
		Wait:       testApp.Wait,
		ErrorChan:  make(chan error),
		MailerChan: make(chan Message, 100),
		DoneChan:   make(chan bool),
	}
}

func listenToTestMailer() {
	for {
		select {
		case <-testApp.Mailer.MailerChan:
		case <-testApp.Mailer.ErrorChan:
		case <-testApp.Mailer.DoneChan:
			return
		}
	}
}

func initLoggers() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func listenToTestLoggers() {
	for {
		select {
		case err := <-testApp.ErrorChan:
			testApp.ErrorLog.Println(err)
		case <-testApp.ErrorChanDone:
			return
		}
	}
}

func getContext(r *http.Request) context.Context {
	ctx, err := testApp.Session.Load(r.Context(), r.Header.Get("X-Session"))

	if err != nil {
		log.Println(err)
	}

	return ctx
}
