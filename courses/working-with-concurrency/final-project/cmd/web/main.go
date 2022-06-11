package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"final_project/data"
)

const (
	dbMaxTries = 10
	dbRetryIn  = 1000 // milliseconds
	webPort    = "8080"
)

var db *sql.DB

func main() {
	db := initDB()

	app := Config{
		Session:       initSession(),
		DB:            db,
		InfoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Wait:          &sync.WaitGroup{},
		Models:        data.New(db),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	app.Mailer = app.createMail()
	go app.listenForMail()

	go app.listenForShutdown()

	go app.listenForErrors()

	app.serve()
}

func (app *Config) serve() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("starting web server")
	err := server.ListenAndServe()

	if err != nil {
		app.ErrorLog.Println("cannot start serving HTTP requests")
		log.Panic(err)
	}
}

// TODO: Singleton with sync.Once()?
func initDB() *sql.DB {
	if db == nil {
		conn, err := connectToDB()

		if err != nil {
			log.Panic("cannot connect to database")
		}

		db = conn
	}

	return db
}

func connectToDB() (*sql.DB, error) {
	tries := 0
	dns := os.Getenv("DSN")

	for {
		conn, err := openDB(dns)

		if tries > dbMaxTries {
			return nil, fmt.Errorf("cannot connect to the database")
		}

		if err != nil {
			log.Println("database not ready yet")
			tries++
			time.Sleep(time.Duration(dbRetryIn) * time.Millisecond)
			continue
		}

		log.Println("connected to database")
		return conn, nil
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initSession() *scs.SessionManager {

	// https://tech.shaadi.com/2021/10/05/serialize-using-gob-in-golang
	// Gob is a serialization form specific to Go data types
	gob.Register(data.User{})

	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
}

func (app *Config) listenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.ErrorChanDone:
			return
		}
	}
}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	app.Wait.Wait()
	app.Mailer.DoneChan <- true
	app.ErrorChanDone <- true
	app.InfoLog.Println("shutting down the application")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
	close(app.ErrorChan)
	close(app.ErrorChanDone)
}

func (app *Config) createMail() Mail {
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	// TODO: Use environment variables
	return Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromAddress: "info@example.com",
		FromName:    "Info",
		Wait:        app.Wait,
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDoneChan,
	}
}
