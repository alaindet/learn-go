package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			panicValue := recover()
			if panicValue != nil {
				w.Header().Set("Connection", "close")
				err := fmt.Errorf("%v", panicValue)
				app.httpErr.InternalServerError(w, r, err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
