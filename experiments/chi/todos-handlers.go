package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func todosRoutes(r chi.Router) {
	r.Get("/", getTodosHandler)
	r.Post("/", createTodoHandler)

	r.Route("/{todoID:[0-9]+}", func(r chi.Router) {
		r.Use(fetchTodoMiddleware)
		r.Get("/", getTodoHandler)
		r.Put("/", updateTodoHandler)
		r.Delete("/", deleteTodoHandler)
	})
}

func fetchTodoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "todoID")
		todo, err := getTodoById(todoID)

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "todo", todo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte("createTodoHandler")
	w.Write(response)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte("getTodosHandler")
	w.Write(response)
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, ok := ctx.Value("todo").(*Todo)

	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	response := []byte(fmt.Sprintf("getTodoHandler: %+v", todo))
	w.Write(response)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, ok := ctx.Value("todo").(*Todo)

	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	response := []byte(fmt.Sprintf("updateTodoHandler: %+v", todo))
	w.Write(response)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, ok := ctx.Value("todo").(*Todo)

	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	response := []byte(fmt.Sprintf("deleteTodoHandler: %+v", todo))
	w.Write(response)
}
