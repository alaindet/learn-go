package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	dbMaxTries = 10
	dbRetryIn  = 1000 // milliseconds
	webPort    = "8080"
)

var db *sql.DB

func main() {
	app := Config{
		Session:  initSession(),
		DB:       initDB(),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Wait:     &sync.WaitGroup{},
	}

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
