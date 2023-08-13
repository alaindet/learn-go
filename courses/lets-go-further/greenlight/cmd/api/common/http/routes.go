package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Get(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodGet, path, handler)
}

func Post(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPost, path, handler)
}

func Put(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPut, path, handler)
}

func Patch(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPatch, path, handler)
}

func Delete(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodDelete, path, handler)
}
