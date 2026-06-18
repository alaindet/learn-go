package main

import (
	data "authentication/data"
	commonJSON "common/json"
	"database/sql"
	"net/http"
)

type App struct {
	DB     *sql.DB
	Models data.Models
}

func (app *App) ReadJSON(
	w http.ResponseWriter,
	r *http.Request,
	toFill any,
) error {
	return commonJSON.ReadRequest(w, r, toFill)
}

func (app *App) WriteJSON(
	w http.ResponseWriter,
	status int,
	jsonData any,
	headers ...http.Header,
) error {
	return commonJSON.WriteResponse(w, status, jsonData, headers...)
}

func (app *App) WriteJSONErr(
	w http.ResponseWriter,
	err error,
	status ...int,
) error {
	return commonJSON.WriteErrorResponse(w, err, status...)
}
