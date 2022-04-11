package gomitolo

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (g *Gomitolo) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	if g.Debug {
		mux.Use(middleware.Logger)
	}

	mux.Use(middleware.Recoverer)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to Gomitolo")
	})

	return mux
}
