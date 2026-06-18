package main

import (
	"common"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	common.WriteJSONResponse(w, http.StatusOK, common.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	})
}
