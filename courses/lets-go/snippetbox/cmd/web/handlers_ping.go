package main

import "net/http"

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
