package main

import (
	commonJSON "common/json"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	commonJSON.WriteResponse(w, http.StatusOK, commonJSON.Response{
		Message: "Hit the broker",
	})
}