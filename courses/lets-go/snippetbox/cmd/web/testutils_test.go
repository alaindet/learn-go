package main

// TODO: Refactor this into a package?
import (
	"bytes"
	"database/sql"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"snippetbox.dev/internal/models/mocks"
)

// Returns an application stub
func newTestApplication(t *testing.T) *application {

	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	return &application{
		errorLog:       log.New(io.Discard, "", 0),
		infoLog:        log.New(io.Discard, "", 0),
		snippets:       mocks.NewSnippetModel(nil),
		users:          mocks.NewUserModel(nil),
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}
}

// Wrap the test server in a struct in order to add methods to it
type testServer struct {
	*httptest.Server
}

// Creates a fake test server
func newTestServer(t *testing.T, handler http.Handler) *testServer {
	server := httptest.NewTLSServer(handler)
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	server.Client().Jar = jar

	// Disable following any 3xx redirect
	server.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{server}
}

// Helper to perform a GET request on the test server
// Returns the status code, the headers and the stringified body
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {

	res, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)
	return res.StatusCode, res.Header, string(body)
}

func (ts *testServer) postForm(
	t *testing.T,
	urlPath string,
	form url.Values,
) (int, http.Header, string) {
	res, err := ts.Client().PostForm(ts.URL+urlPath, form)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, res.Header, string(body)
}

func newTestDB(t *testing.T) *sql.DB {

	// TODO: Move?
	var (
		username = "snippetboxtest"
		password = "snippetboxtest"
		host     = "localhost"
		port     = "5432"
		dbname   = "snippetboxtest"
	)

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		t.Fatal(err)
	}

	// Try reaching for the database
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	// Database setup
	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	// Database teardown
	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	})

	return db
}

var csrfTokenRegex = regexp.MustCompile(
	`<input type="hidden" name="csrf_token" value="(.+)">`,
)

func extractCSRFToken(t *testing.T, body string) string {
	matches := csrfTokenRegex.FindStringSubmatch(body)

	// There must be at least 2 matches as
	// the first match is always the whole string
	if len(matches) < 2 {
		t.Fatal("no CSRF token found in body")
	}

	csrfToken := matches[1]

	return html.UnescapeString(csrfToken)
}
