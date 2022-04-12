package gomitolo

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (g *Gomitolo) NewRouter() *chi.Mux {
	mux := chi.NewRouter()

	// Global middleware
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	// Testing middleware
	if g.Debug {
		mux.Use(middleware.Logger)
	}

	// Global error middleware
	mux.Use(middleware.Recoverer)

	return mux
}
