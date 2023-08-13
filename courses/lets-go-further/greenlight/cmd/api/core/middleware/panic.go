package middleware

import (
	"fmt"
	"greenlight/cmd/api/core"
	"net/http"
)

func RecoverPanic(app *core.Application, next http.Handler) http.Handler {

	// No setup

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.InternalServerErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
