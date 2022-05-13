package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Todo App v12345")
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	todos := db.GetTodos()
	json.NewEncoder(w).Encode(todos)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	routeParams := mux.Vars(r)
	todoId := routeParams["todoid"]
	todo, err := db.GetTodoById(todoId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Todo not found"}`)) // TODO
		return
	}

	json.NewEncoder(w).Encode(todo)
}
